package main

import (
	"github.com/nntaoli-project/goex/v2/binance"
	"github.com/nntaoli-project/goex/v2/httpcli"
	"github.com/nntaoli-project/goex/v2/huobi"
	"github.com/nntaoli-project/goex/v2/logger"
	"github.com/nntaoli-project/goex/v2/model"
	"github.com/nntaoli-project/goex/v2/okx"
	"github.com/nntaoli-project/goex/v2/options"
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
	logger.SetLevel(logger.INFO)
	DefaultHttpCli.SetTimeout(5)
	_, _, err := OKx.Spot.GetExchangeInfo()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	btcUSDTCurrencyPair, err := OKx.Spot.NewCurrencyPair(model.BTC, model.USDT)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	okxPrvApi := OKx.Spot.NewPrvApi(
		options.WithApiKey(""),
		options.WithApiSecretKey(""),
		options.WithPassphrase(""))

	list, _, err := okxPrvApi.GetKline(btcUSDTCurrencyPair, model.Kline_5min)
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println(list[0])
	}
}
