package bitwise

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func CollectFromURLAndSearch(url string, search string) (result types.Result) {
	// Create a new collector
	c := colly.NewCollector()

	// Find and visit the target URL
	c.OnHTML("div[class*='layout-base']", func(e *colly.HTMLElement) {
		// Check if the div contains the desired text
		if strings.Contains(e.Text, search) {
			// Look for the div containing the value
			e.ForEach("div", func(_ int, el *colly.HTMLElement) {
				if strings.Contains(el.Text, search) {
					// Get the next div element which contains the figure
					figure := el.DOM.Next().Text()
					// Print the figure
					inputClean := strings.ReplaceAll(figure, ",", "")
					totalAssetInTrust, _ := strconv.ParseFloat(inputClean, 64)
					result.TotalAsset = totalAssetInTrust
					return
				}
			})
		}
	})

	// Visit the website
	c.Visit(url)

	c.Wait()

	return result
}
