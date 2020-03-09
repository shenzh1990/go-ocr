package carry

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/fcoin"
	"net"
	"net/http"
	"net/url"
	"time"
)

var fcws = fcoin.NewFCoinWs(&http.Client{
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
})

func init() {
	fcws.ProxyUrl("http://127.0.0.1:1088")
	fcws.SetCallbacks(PrintfTicker, PrintfDepth, PrintfTrade, PrintfKline)
}

func GetDepthWithWs() {
	//return
	fcws.SubscribeDepth(goex.BTC_USDT, 20)
	time.Sleep(time.Second * 10)
}
func GetTickerWithWs() {
	//return
	fcws.SubscribeTicker(goex.BTC_USDT)
	time.Sleep(time.Second * 10)
}
