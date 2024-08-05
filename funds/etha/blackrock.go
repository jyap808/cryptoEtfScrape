package etha

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/blackrock"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result, err error) {
	url := "https://www.blackrock.com/us/individual/products/337614/fund/1464253357814.ajax?tab=all&fileType=json"
	return blackrock.CollectFromURLAndTicker(url, "ETH")
}
