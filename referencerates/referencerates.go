package referencerates

import (
	"github.com/jyap808/ethEtfScrape/referencerates/cmeny"
)

func CMENYCollect() (cmeny.ReferenceRates, error) {
	return cmeny.GetReferenceRates()
}
