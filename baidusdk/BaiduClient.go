package baidusdk

import (
	"BitCoin/pkg/settings"
	"errors"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type BaiduMapClient struct {
	Url     string
	SK      string
	Geocode GeoCode
}

var BMClient *BaiduMapClient

//初始化
func init() {
	BMClient = &BaiduMapClient{
		Url: settings.BitConfig.Baidu.BaiduMap.Url,
		SK:  settings.BitConfig.Baidu.BaiduMap.SK,
	}
}

func (d *BaiduMapClient) GetGeoCode(address string) (string, error) {

	Url := d.Url + "/geocoding/v3/?address=" + address + "&output=json&ak=" + d.SK
	_, body, errs := gorequest.New().Get(Url).EndBytes()
	if len(errs) > 0 {
		return "", errors.New("access get error")
	}
	if jsoniter.Get(body, "status").ToInt32() != 0 {
		return "", errors.New(string(body))
	}
	return string(body), nil
}
