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
{"imagetype":"0","imagebase":"","whitelist":"","languages":"chi_sim","trim":"","uuid":"1234"}
*/
func GetOrderInfo(c *gin.Context) {
	var imagebody services.OcrBody
	err := c.BindJSON(&imagebody)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, err.Error(), ""))
		return
	}
	if imagebody.Unid == "" {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, "请输入Unid", ""))
		return
	}
	bT := time.Now() // 开始时间
	text, err := services.ImageOcr(imagebody)
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("ImageOcr Run time: ", eT)
	if err != nil {
		c.String(http.StatusOK, common.JsonResponse(e.ERROR, err.Error(), ""))
		return
	}
	bT = time.Now() // 开始时间
	pts := services.Participle(text)
	eT = time.Since(bT) // 从开始到当前所消耗的时间
	cmn.Info("Participle Run time: ", eT)
	order := Order{}
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
	order.OcrText = text
	cmn.Info(order)
	redisutil.RDS.SetEx(redisutil.DEFAULT_REDIS_PRE_KEY+imagebody.Unid, order, 1000)
	c.String(http.StatusOK, common.JsonResponse(e.SUCCESS, "", order))
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
