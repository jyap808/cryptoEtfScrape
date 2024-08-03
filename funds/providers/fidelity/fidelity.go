package fidelity

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/jyap808/cryptoEtfScrape/types"

	"github.com/ledongthuc/pdf"
)

type PDFCoordinates struct {
	Row    RowCoordinates
	Column ColumnCoordinates
}

type RowCoordinates struct {
	Index        int
	ContentIndex int
}

type ColumnCoordinates struct {
	Index        int
	ContentIndex int
}

func CollectFromURLsAndPDFCoordinates(pdfBaseURL string, prospectusURL string, pdfCoordinates PDFCoordinates) (result types.Result) {
	actionExchangeRepositoryURL := getActionExchangeRepositoryURL(prospectusURL)
	collectionID := getCollectionID(actionExchangeRepositoryURL)

	url := fmt.Sprintf(pdfBaseURL, collectionID)

	// Fetch the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error performing request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	// Create a reader from the byte slice
	reader := bytes.NewReader(body)

	r, err := pdf.NewReader(reader, int64(len(body)))
	if err != nil {
		log.Println("Error creating reader:", err)
		return
	}

	pageIndex := 1
	p := r.Page(pageIndex)

	rows, _ := p.GetTextByRow()
	rowHoldings := rows[pdfCoordinates.Row.Index].Content[pdfCoordinates.Row.ContentIndex].S

	cols, _ := p.GetTextByColumn()
	colHoldings := cols[pdfCoordinates.Column.Index].Content[pdfCoordinates.Column.ContentIndex].S

	if rowHoldings == colHoldings {
		total, _ := strconv.ParseFloat(rowHoldings, 64)
		result.TotalAsset = total
	} else {
		log.Println("Fidelity PDF row and column collection mismatch")
	}

	return
}

// The URL redirects to www.actionsxchangerepository.fidelity.com
func getActionExchangeRepositoryURL(url string) (redirectURL string) {
	c := colly.NewCollector()

	// Find and extract the redirect URL
	c.OnHTML("a", func(e *colly.HTMLElement) {
		redirectURL = e.Attr("href")
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	return redirectURL
}

func getCollectionID(url string) (collectionID int) {
	c := colly.NewCollector()

	c.OnHTML("td", func(e *colly.HTMLElement) {
		// Check if the fundDocumentType is "DALY"
		if strings.Contains(e.Attr("onclick"), "'DALY'") {
			// Extract
			collectionID = extractCollectionIDFromOnClick(e.Attr("onclick"))
		}
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		log.Println("Error:", err)
		return 0
	}

	return collectionID
}

func extractCollectionIDFromOnClick(onClick string) int {
	// Split the onClick attribute by comma and extract the element
	parts := strings.Split(onClick, ",")
	if len(parts) >= 5 {
		// Remove surrounding quotes and trim whitespace
		rawID := strings.TrimSpace(strings.Trim(parts[6], "'"))
		ID, _ := strconv.Atoi(rawID)
		return ID
	}
	return 0
}
