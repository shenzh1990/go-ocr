package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-ocr/api/v1"
	"go-ocr/middleware/cors"
	"go-ocr/pkg/settings"
)

type IgnoreGinStdoutWritter struct{}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Cors())
	gin.DisableConsoleColor() // 关闭Gin的日志颜色
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.GET("/getSegmentWord", controller.GetSegmentWord)
	r.POST("/get_order_info", v1.GetOrderInfo)
	//r.POST("/login", controller.Login)
	return r
}
