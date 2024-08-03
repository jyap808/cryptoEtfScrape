package ibit

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/blackrock"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result) {
	url := "https://blackrock.com/us/financial-professionals/products/333011/fund/1500962885783.ajax?tab=all&fileType=json"
	return blackrock.CollectFromURLAndTicker(url, "BTC")
}
