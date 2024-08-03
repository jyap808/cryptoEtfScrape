package arkb

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/twentyoneshares"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result) {
	url := "https://cdn.21shares.com/uploads/fund-documents/bny-bank/holdings/product/current/ARKB-Export.csv"
	return twentyoneshares.CollectFromURL(url)
}
