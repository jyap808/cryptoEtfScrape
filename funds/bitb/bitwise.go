package bitb

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/bitwise"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result, err error) {
	return bitwise.CollectFromURLAndSearch("https://bitbetf.com/", "Bitcoin in Trust")
}
