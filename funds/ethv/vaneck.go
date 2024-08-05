package ethv

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/vaneck"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result, err error) {
	url := "https://www.vaneck.com/Main/NavInformationBlock/GetContent/?blockid=280232&ticker=ETHV"
	return vaneck.CollectFromURLAndSearch(url, "Ether in Trust")
}
