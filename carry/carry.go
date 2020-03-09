package carry

import (
	"BitCoin/pkg/settings"
	"fmt"
	"github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/fcoin"
	"net"
	"net/http"
	"net/url"
	"time"
)

var ft = fcoin.NewFCoin(&http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse("http://127.0.0.1:1088")
			return nil, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}, settings.BitConfig.Fcoin.FcoinKey, settings.BitConfig.Fcoin.FcoinSecret)

func PrintfTicker(ticker *goex.Ticker) {
	fmt.Println(ticker)
	////获取价格
	//acc, err := ft.GetAssets();
	//if (err != nil) {
	//	fmt.Println("获取账单信息出错！")
	//}
	//for _,asset := range acc{
	//	if(asset.Currency==goex.EOS){
	//		fmt.Println(asset)
	//	}
	//}
	//fmt.Println(acc)
	//ft.LimitBuy("1",goex.FloatToString(ticker.Buy,2),ticker.Pair)

}
func PrintfDepth(depth *goex.Depth) {
	//acc, err := ft.GetAssets();
	//if (err != nil) {
	//	fmt.Println("获取账单信息出错！")
	//}
	//for _,asset := range acc{
	//	if(asset.Currency==goex.EOS){
	//		fmt.Println(asset)
	//	}
	//}
	fmt.Println(depth)

}
func PrintfTrade(trade *goex.Trade) {
	fmt.Println(trade)
}
func PrintfKline(kline *goex.Kline, period int) {
	fmt.Println(kline)
}
