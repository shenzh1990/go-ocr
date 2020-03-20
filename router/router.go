package router

import (
	"BitCoin/api/blog_v1"
	"BitCoin/controller"
	"BitCoin/middleware/cors"
	"BitCoin/middleware/jwt"
	"BitCoin/pkg/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Cors())
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getSegmentWord", controller.GetSegmentWord)
	r.POST("/getSegment", controller.GetSegment)
	r.POST("/login", controller.Login)
	api_blog_v1 := r.Group("/api/v1")
	api_blog_v1.Use(jwt.AuthorizedMiddelware(settings.JwtSecret))
	{
		//获取标签列表
		api_blog_v1.GET("/tags", blog_v1.GetTags)
		//新建标签
		api_blog_v1.POST("/tags", blog_v1.AddTag)
		//更新指定标签
		api_blog_v1.PUT("/tags/:id", blog_v1.EditTag)
		//删除指定标签
		api_blog_v1.DELETE("/tags/:id", blog_v1.DeleteTag)
	}
	return r
}
