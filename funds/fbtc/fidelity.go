package fbtc

import (
	"github.com/jyap808/cryptoEtfScrape/funds/providers/fidelity"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func Collect() (result types.Result) {
	pdfBaseURL := "https://www.actionsxchangerepository.fidelity.com/ShowDocument/documentPDF.htm?clientId=Fidelity&applicationId=MFL&securityId=315948109&docType=DALY&docFormat=pdf&securityIdType=CUSIP&collectionId=%d&docName=1.WOB-DALY.pdf&criticalIndicator=N&pdfReaderStatus=Y"
	prospectusURL := "https://fundresearch.fidelity.com/prospectus/eproredirect?clientId=Fidelity&applicationId=MFL&securityIdType=CUSIP&critical=N&securityId=315948109"
	pdfCoordinates := fidelity.PDFCoordinates{
		Row:    fidelity.RowCoordinates{Index: 0, ContentIndex: 9},
		Column: fidelity.ColumnCoordinates{Index: 1, ContentIndex: 18},
	}

	return fidelity.CollectFromURLsAndPDFCoordinates(pdfBaseURL, prospectusURL, pdfCoordinates)
}
