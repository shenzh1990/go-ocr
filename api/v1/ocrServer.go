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

	OCR_SUCCESS string = "1"
	OCR_ING     string = "2"
	OCR_FAIL    string = "3"
	OCR_UNKNOWN string = "0"
)

type Order struct {
	OcrStatue string `json:"ocr_statue"` // 0:未识别 1:识别成功 2:识别中 3:识别失败
	OcrMsg    string `json:"ocr_msg"`    // 识别失败的原因
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
	//先查询是否存在，存在就直接返回，不存在就执行ocr
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
			result, _ := redisutil.RDS.SetNxEx(redisutil.DEFAULT_REDIS_PRE_NX_KEY+imagebody.Unid, imagebody.Unid, 1000)
			defer func() {
				redisutil.RDS.Del(redisutil.DEFAULT_REDIS_PRE_NX_KEY + imagebody.Unid)
			}()
			flag, ok := result.(int64)
			if ok {
				//1 不存在并设置值
				//0 存在失败
				if flag == 1 {
					order := Order{}
					order.OcrStatue = OCR_ING
					//设置执行中
					err = redisutil.RDS.SetEx(redisutil.DEFAULT_REDIS_PRE_KEY+imagebody.Unid, common.JsonResponse(e.SUCCESS, "", order), 2592000)
					//执行异步识别
					order = GetOrderSync(imagebody)
					//设置执行状态
					err = redisutil.RDS.SetEx(redisutil.DEFAULT_REDIS_PRE_KEY+imagebody.Unid, common.JsonResponse(e.SUCCESS, "", order), 2592000)
				}
			}
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
func GetOrderSync(imagebody services.OcrBody) Order {
	order := Order{}
	bT := time.Now() // 开始时间
	text, err := services.ImageOcr(imagebody)
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("ImageOcr Run time: ", eT)
	if err != nil {
		order.OcrStatue = OCR_FAIL
		order.OcrMsg = "ocr失败：" + err.Error()
		return order
	}
	bT = time.Now() // 开始时间
	pts := services.Participle(text)
	eT = time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("Participle Run time: ", eT)

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
		order.OcrStatue = OCR_FAIL
		order.OcrMsg = "ocr失败：无法正确分词"
		order.OcrText = text
		return order
	}
	order.OcrStatue = OCR_SUCCESS
	order.OcrMsg = "ocr成功"
	order.OcrText = text
	return order
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
