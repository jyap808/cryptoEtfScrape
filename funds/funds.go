package funds

import (
	"github.com/jyap808/ethEtfScrape/funds/eth"
	"github.com/jyap808/ethEtfScrape/funds/ethe"
	"github.com/jyap808/ethEtfScrape/funds/ethv"
	"github.com/jyap808/ethEtfScrape/funds/ethw"
	"github.com/jyap808/ethEtfScrape/types"
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
