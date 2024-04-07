package services

import (
	"encoding/base64"
	"github.com/otiai10/gosseract/v2"
	"go-ocr/participle"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type OcrBody struct {
	//默认0 imagebase 为url
	//1 imagebase 为base64
	ImageType string `json:"imagetype"`
	ImageBase string `json:"imagebase"`
	Trim      string `json:"trim"`
	//默认中文chi_sim
	//英文 eng
	Languages string `json:"languages"`
	Whitelist string `json:"whitelist"`
	//必填
	Unid string `json:"Unid"`
}

func Participle(text string) []participle.TextWord {
	//处理没用的信息
	//text = "14:14                            2 公\nN   AN 中二            万   从\n《            订单详情          (CO\n Biaze(毕亚兹)京东自营旗舰店\n       EREID 毕亚效 高速USB3.0数据线..， *15.9\n人 1 数量 x1 3.0延长线【公对公镀金]-0.5米，[Ev\n伟力  支持7天无理由退货                 ，\nRE  从生 ，，\n卖了换钱 ) (加购物车 ) | 申请售后 )\n实付款                     共减半6.2 合计闻15.7》\n订单编号                    291261877823 复制\n支付方式                                        白条 》\n支付时间                                    2024-04-04 23:14:54\n下单时间                                    2024-04-04 23:14:50\n配送方式                            京东快递\n期望配送时间              2024-04-05,15:00-21:00\n收起 ^\n服务中心                                >\n价格保护                   去申请 ”只换不修                    #18起，\n7天价保              性能故障，只换不修\n导斑兴\n快速解决问题\n怎么申请售后      怎么查看我的发票     更多\n和 店铺客服            京东客服\n四 商品/活动             是 物流/售后/平台\n更多          查看发票 ) | 退换/售后 | 人知E多MXEE\n本"
	text = strings.Replace(text, "\\n", "\n", -1)
	text = strings.Replace(text, " ", "", -1)
	segments := participle.PartSem.Segment([]byte(text))
	return participle.SegmentsToObject(segments, false, "")
}

func ImageOcr(imageBody OcrBody) (string, error) {
	//return "", nil
	//创建一个临时文件
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	if err != nil {
		return "", err
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()
	if imageBody.ImageType == "1" {
		s := regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(imageBody.ImageBase, "")
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return "", err
		}
		tempfile.Write(b)
	} else {
		// 下载图片到临时文件
		resp, err := http.Get(imageBody.ImageBase)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		// 将响应体写入临时文件
		_, err = io.Copy(tempfile, resp.Body)
		if err != nil {
			return "", err
		}
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.Languages = []string{"chi_sim"}
	if imageBody.Languages != "" {
		client.Languages = strings.Split(imageBody.Languages, ",")
	}
	client.SetImage(tempfile.Name())
	if imageBody.Whitelist != "" {
		client.SetWhitelist(imageBody.Whitelist)
	}

	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return strings.Trim(text, imageBody.Trim), nil
}
