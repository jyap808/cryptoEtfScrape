package funds

import (
	"github.com/jyap808/cryptoEtfScrape/funds/eth"
	"github.com/jyap808/cryptoEtfScrape/funds/ethe"
	"github.com/jyap808/cryptoEtfScrape/funds/ethv"
	"github.com/jyap808/cryptoEtfScrape/funds/ethw"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func EthCollect() types.Result {
	return eth.Collect()
}

func EtheCollect() types.Result {
	return ethe.Collect()
}

func EthvCollect() types.Result {
	return ethv.Collect()
}

func EthwCollect() types.Result {
	return ethw.Collect()
}
