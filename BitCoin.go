package main

import (
	"BitCoin/controller"
	"BitCoin/pkg/settings"
	"BitCoin/router"
	"fmt"
	"github.com/gin-gonic/gin"
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
