package main

import (
	"github.com/nntaoli-project/goex/v2/binance"
	"github.com/nntaoli-project/goex/v2/httpcli"
	"github.com/nntaoli-project/goex/v2/huobi"
	"github.com/nntaoli-project/goex/v2/logger"
	"github.com/nntaoli-project/goex/v2/model"
	"github.com/nntaoli-project/goex/v2/okx"
	"log"
	"reflect"
)

var (
	DefaultHttpCli = httpcli.Cli
)

var (
	OKx     = okx.New()
	Binance = binance.New()
	HuoBi   = huobi.New()
)

func SetDefaultHttpCli(cli httpcli.IHttpClient) {
	logger.Infof("use new http client implement: %s", reflect.TypeOf(cli).Elem().String())
	httpcli.Cli = cli
}

func main() {
	logger.SetLevel(logger.DEBUG)
	DefaultHttpCli.SetTimeout(10)

	_, _, err := OKx.Spot.GetExchangeInfo() //建议调用
	if err != nil {
		panic(err)
	}
	btcUSDTCurrencyPair, err := OKx.Spot.NewCurrencyPair(model.BTC, model.USDT)
	if err != nil {
		panic(err)
	}
	log.Println(OKx.Spot.GetTicker(btcUSDTCurrencyPair))

	okxPrvApi := OKx.Spot.NewPrvApi()
	okxPrvApi.GetKline(btcUSDTCurrencyPair, model.Kline_1min)
	order, _, err := okxPrvApi.CreateOrder(btcUSDTCurrencyPair, 0.01, 18000, model.Spot_Buy, model.OrderType_Limit)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(order)
	}
}
