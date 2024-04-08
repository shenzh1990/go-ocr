package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/common"
	"go-ocr/participle"
	"go-ocr/pkg/e"
	"go-ocr/redisutil"
	"go-ocr/services"
	"net/http"
	"time"
)

const (
	POS_YFK  string = "yfk"
	POS_DDBH string = "ddbh"
	POS_XDSJ string = "xdsj"
)

type Order struct {
	Dues      string `json:"dues"`
	OrderTime string `json:"order_time"`
	OrderNum  string `json:"order_num"`
	OcrText   string `json:"ocr_txt"`
}

// c *gin.Context
/**
{"imagetype":"0","imagebase":"","whitelist":"","languages":"chi_sim","trim":"","unid":"1234"}
*/
func OcrOrderInfo(c *gin.Context) {
	var imagebody services.OcrBody
	err := c.BindJSON(&imagebody)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, err.Error(), ""))
		return
	}
	if imagebody.Unid == "" || imagebody.ImageBase == "" {
		c.String(http.StatusOK, common.JsonResponse(e.INVALID_PARAMS, "请输入确认输入Unid和Image", ""))
		return
	}
	result, err := redisutil.RDS.Get(redisutil.DEFAULT_REDIS_PRE_KEY + imagebody.Unid)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, "redis 获取失败", ""))
		return
	}
	byteValue, ok := result.([]byte)
	// 检查转换是否成功
	if ok {
		strValue := string(byteValue)
		if strValue != "" {
			c.String(http.StatusOK, strValue)
			return
		}
	} else {
		go func() {
			err = redisutil.RDS.SetEx(redisutil.DEFAULT_REDIS_PRE_KEY+imagebody.Unid, GetOrderSync(imagebody), 1000)
			cmn.Error(err)
		}()
	}
	c.String(http.StatusOK, common.JsonResponse(e.SUCCESS, "", ""))
}

/*
*
{"unid":"1234"}
*/
func GetOrderInfo(c *gin.Context) {
	var imagebody services.OcrBody
	err := c.BindJSON(&imagebody)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, err.Error(), ""))
		return
	}
	if imagebody.Unid == "" {
		c.String(http.StatusOK, common.JsonResponse(e.INVALID_PARAMS, "请输入Unid", ""))
		return
	}

	result, err := redisutil.RDS.Get(redisutil.DEFAULT_REDIS_PRE_KEY + imagebody.Unid)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, "redis 获取失败", ""))
		return
	}
	strValue, ok := result.([]byte)
	// 检查转换是否成功
	if !ok {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, "未获得值", result))
		return
	}
	c.String(http.StatusOK, string(strValue))
}
func GetOrderSync(imagebody services.OcrBody) string {
	var Result string
	bT := time.Now() // 开始时间
	text, err := services.ImageOcr(imagebody)
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("ImageOcr Run time: ", eT)
	if err != nil {
		Result = common.JsonResponse(e.ERROR_OCR_IMG, err.Error(), "")
		return Result
	}
	bT = time.Now() // 开始时间
	pts := services.Participle(text)
	eT = time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("Participle Run time: ", eT)
	order := Order{}
	if pts != nil && len(pts) > 0 {
		for i, pt := range pts {
			{
				if pt.Pos == POS_DDBH {
					order.OrderNum = pts[i+1].Text
				}
				if pt.Pos == POS_XDSJ {
					timeStr := ""
					for j := i + 1; j < len(pts); j++ {
						if pts[j].Pos == participle.POS_X {
							if cmn.IsNumber(pts[j].Text) || pts[j].Text == ":" {
								timeStr = timeStr + pts[j].Text
							}
						} else {
							break
						}
					}
					t, err := parseTime(timeStr)
					if err != nil {
						order.OrderTime = timeStr
					} else {
						order.OrderTime = t.Format("2006-01-02 15:04:05")
					}

				}
				if pt.Pos == POS_YFK {
					moneyStr := ""
					for j := i + 1; j < len(pts); j++ {
						if pts[j].Pos == participle.POS_X {
							if cmn.IsNumber(pts[j].Text) || pts[j].Text == "." {
								moneyStr = moneyStr + pts[j].Text
							}
						} else if pts[j].Pos != participle.POS_X && moneyStr != "" {
							break
						}
					}
					order.Dues = moneyStr
				}
			}
		}
	} else {
		Result = common.JsonResponse(e.ERROR_OCR_IMG, "没有正确的ocr数据", "")
		return Result
	}
	order.OcrText = text
	Result = common.JsonResponse(e.SUCCESS, "", order)
	return Result
}
func parseTime(str string) (time.Time, error) {
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:0405",
		"2006-01-02 1504:05",
		"2006-01-0215:04:05",
		"2006-0102 15:04:05",
		"200601-02 15:04:05",
		"2006-01-02 150405",
		"2006-01-0215:0405",
		"2006-0102 15:0405",
		"200601-02 15:0405",
		"2006-01-021504:05",
		"2006-0102 1504:05",
		"200601-02 1504:05",
		"2006-010215:04:05",
		"200601-0215:04:05",
		"20060102 15:04:05",
		"2006-01-02150405",
		"2006-0102 150405",
		"200601-02 150405",
		"2006-010215:0405",
		"20060102 15:0405",
		"2006-01021504:05",
		"200601-021504:05",
		"2006010215:04:05",
		"2006-0102150405",
		"2006010215:0405",
		"200601021504:05",
		"20060102150405",
	}

	for _, format := range formats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("无法解析的时间格式: %s", str)
}
