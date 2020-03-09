package common

import (
	"BitCoin/pkg/settings"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func JsonResponse(code int, data interface{}) string {
	response := make(map[string]interface{})
	response["code"] = code
	response["data"] = data
	js, err := json.Marshal(response)
	if err != nil {
		return err.Error()
	}
	return string(js)
}
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}

	return result
}
