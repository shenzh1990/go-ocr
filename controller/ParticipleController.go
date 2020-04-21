package controller

import (
	. "BitCoin/common"
	"BitCoin/model"
	"BitCoin/participle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

//获得自定义原词
func GetSegmentWord(c *gin.Context) {

	text := c.DefaultQuery("text", "高血压")
	rule := model.UserRule{
		Id:           0,
		Text:         text,
		Frequency:    0,
		PartOfSpeech: "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
	rs := rule.GetUserRules()
	c.String(http.StatusOK, JsonResponse(1, rs))
	return
}

//post
//body json {"text":"value1",
//            "seachmode":"1",
//             "ps","nil"
// }
// text 分词文本
// seachmode 分词模式  1 表示搜索模式  0表示普通模式 默认0
// ps 词性 获取指定词性的词 ,""表示不处理，
func GetSegment(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]string
	if err := json.Unmarshal(data, &dat); err == nil {
		seachmode := false

		if dat["text"] == "" {
			c.String(http.StatusOK, JsonResponse(0, "请提供需要分词的文字"))
			return
		}
		if dat["seachmode"] == "1" {
			seachmode = true
		}

		segments := participle.PartSem.Segment([]byte(dat["text"]))
		c.String(http.StatusOK, JsonResponse(0, participle.SegmentsToObject(segments, seachmode, dat["ps"])))
		return
	}
	return
}
