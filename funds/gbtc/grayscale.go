package gbtc

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/grayscale"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result) {
	return grayscale.CollectFromURL("https://etfs.grayscale.com/gbtc")
}
