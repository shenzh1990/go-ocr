package controller

import (
	. "BitCoin/common"
	. "BitCoin/model"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := Customer{
		Username: username,
		Password: password,
	}
	c.String(http.StatusOK, user.Register())
}
func Login(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]string
	if err := json.Unmarshal(data, &dat); err == nil {
		user := Customer{}
		user = user.SelectUserByName(dat["username"])
		if dat["username"] == "" || dat["password"] == "" {
			c.String(http.StatusOK, JsonResponse(0, "账户名密码不能为空"))
			return
		}
		if dat["password"] == user.Password {
			c.String(http.StatusOK, JsonResponse(1, gin.H{"token": setToken(dat["username"])}))
			return
		} else {
			c.String(http.StatusOK, JsonResponse(0, "用户名密码不正确"))
			return
		}
	}

}

func setToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user"] = username
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("mobile"))
	if err != nil {
		return ""
	}
	return tokenString
}
