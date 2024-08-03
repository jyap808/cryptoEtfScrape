package hodl

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/vaneck"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result) {
	url := "https://www.vaneck.com/Main/NavInformationBlock/GetContent/?blockid=252190&ticker=HODL"
	return vaneck.CollectFromURLAndSearch(url, "Bitcoin in Trust")
}
