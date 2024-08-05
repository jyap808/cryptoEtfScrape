package ezbc

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/franklin"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result, err error) {
	return franklin.CollectWithFundIDAndSearch(39639, "BITCOIN")
}
