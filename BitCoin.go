package main

import (
	"BitCoin/controller"
	"BitCoin/pkg/settings"
	"BitCoin/router"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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

	authorized := r.Group("/user", AuthorizedMiddelware())

	authorized.POST("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})

}
func AuthorizedMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("mobile"), nil
			})
		if err == nil {
			if token.Valid {
				fmt.Println(token.Claims)
				c.Next()
			} else {
				c.String(http.StatusUnauthorized, "Token is not valid")
			}
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized access to this resource")
		}
	}
}
