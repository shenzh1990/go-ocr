package main

import (
	"BitCoin/controller"
	"BitCoin/pkg/settings"
	"BitCoin/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        r,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	r.POST("/login", controller.Login)
	//r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

}

func sendSms() {
	//url:="http://223.82.247.212:8080/sjb/HttpSendSMSService"
	url := "https://jxfywx.eheren.com/sjb/HttpSendSMSService"
	content := "<?xml version='1.0' encoding='utf-8'?> <svc_init ver='2.0.0'> <sms ver='2.0.0'> <client> <id>79</id> <pwd>EnWJKIIYKTVDrIm691wt2dit28L1lsNA595Zbba7HNs8Zbeya8bexQ==</pwd> <serviceid>106575360441075</serviceid> </client> <sms_info> <phone>15158133528      </phone> <content>测试11</content> </sms_info> </sms> </svc_init>"
	_, body, errs := gorequest.New().Post(url).Set("Content-Type", "application/xml").
		SendString(content).End()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	fmt.Println("入参：" + content)
	fmt.Print("出参：" + string(body))
}
