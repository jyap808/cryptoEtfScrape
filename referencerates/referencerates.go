package referencerates

import (
	"github.com/jyap808/cryptoEtfScrape/referencerates/cmeny"
)

func CMENYCollect() (cmeny.ReferenceRates, error) {
	return cmeny.GetReferenceRates()
}
