package funds

import (
	// BTC
	"github.com/jyap808/cryptoEtfScrape/funds/arkb"
	"github.com/jyap808/cryptoEtfScrape/funds/bitb"
	"github.com/jyap808/cryptoEtfScrape/funds/brrr"
	"github.com/jyap808/cryptoEtfScrape/funds/btcw"
	"github.com/jyap808/cryptoEtfScrape/funds/defi"
	"github.com/jyap808/cryptoEtfScrape/funds/ezbc"
	"github.com/jyap808/cryptoEtfScrape/funds/fbtc"
	"github.com/jyap808/cryptoEtfScrape/funds/gbtc"
	"github.com/jyap808/cryptoEtfScrape/funds/hodl"
	"github.com/jyap808/cryptoEtfScrape/funds/ibit"

	// ETH
	"github.com/jyap808/cryptoEtfScrape/funds/ceth"
	"github.com/jyap808/cryptoEtfScrape/funds/eth"
	"github.com/jyap808/cryptoEtfScrape/funds/etha"
	"github.com/jyap808/cryptoEtfScrape/funds/ethe"
	"github.com/jyap808/cryptoEtfScrape/funds/ethv"
	"github.com/jyap808/cryptoEtfScrape/funds/ethw"
	"github.com/jyap808/cryptoEtfScrape/funds/ezet"
	"github.com/jyap808/cryptoEtfScrape/funds/feth"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collector(ticker string) types.Result {
	switch ticker {
	// BTC
	case "ARKB":
		return arkb.Collect()
	case "BITB":
		return bitb.Collect()
	case "BRRR":
		return brrr.Collect()
	case "BTCW":
		return btcw.Collect()
	case "DEFI":
		return defi.Collect()
	case "EZBC":
		return ezbc.Collect()
	case "FBTC":
		return fbtc.Collect()
	case "GBTC":
		return gbtc.Collect()
	case "HODL":
		return hodl.Collect()
	case "IBIT":
		return ibit.Collect()
	// ETH
	case "CETH":
		return ceth.Collect()
	case "ETH":
		return eth.Collect()
	case "ETHA":
		return etha.Collect()
	case "ETHE":
		return ethe.Collect()
	case "ETHV":
		return ethv.Collect()
	case "ETHW":
		return ethw.Collect()
	case "EZET":
		return ezet.Collect()
	case "FETH":
		return feth.Collect()
	default:
		return types.Result{}
	}
}
