package ethw

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/bitwise"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result, err error) {
	return bitwise.CollectFromURLAndSearch("https://ethwetf.com/", "ETH in Trust")
}
