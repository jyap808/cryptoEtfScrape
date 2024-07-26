package funds

import (
	"github.com/jyap808/cryptoEtfScrape/funds/ceth"
	"github.com/jyap808/cryptoEtfScrape/funds/eth"
	"github.com/jyap808/cryptoEtfScrape/funds/ethe"
	"github.com/jyap808/cryptoEtfScrape/funds/ethv"
	"github.com/jyap808/cryptoEtfScrape/funds/ethw"
	"github.com/jyap808/cryptoEtfScrape/funds/ezet"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collector(ticker string) types.Result {
	switch ticker {
	case "CETH":
		return ceth.Collect()
	case "ETH":
		return eth.Collect()
	case "ETHE":
		return ethe.Collect()
	case "ETHV":
		return ethv.Collect()
	case "ETHW":
		return ethw.Collect()
	case "EZET":
		return ezet.Collect()
	default:
		return types.Result{}
	}
}
