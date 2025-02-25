package grayscale

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/jyap808/cryptoEtfScrape/types"
)

func CollectFromURL(url string) (result types.Result, err error) {
	c := colly.NewCollector()

	// Tracking flag
	dataFound := false

	c.OnHTML("script", func(e *colly.HTMLElement) {
		if dataFound {
			return
		}

		scriptContent := e.Text

		// Skip scripts that don't contain next_f
		if !strings.Contains(scriptContent, "self.__next_f.push") {
			return
		}

		// Find the pricingData section which should contain our target fields
		if idx := strings.Index(scriptContent, "pricingData"); idx != -1 {
			// Extract a reasonable chunk of text after "pricingData"
			endIdx := idx + 2000
			if endIdx > len(scriptContent) {
				endIdx = len(scriptContent)
			}
			relevantSection := scriptContent[idx:endIdx]

			// Now look for both target fields in this focused section
			if strings.Contains(relevantSection, "totalAssetInTrust") &&
				strings.Contains(relevantSection, "pricingDataDate") {

				extractedResult, extractErr := extractDataFromScript(relevantSection)
				if extractErr == nil {
					result = extractedResult
					dataFound = true
					return
				}
				log.Printf("Found matching section but extraction failed: %v", extractErr)
			}
		}
	})

	// visiting the target page
	c.Visit(url)

	c.Wait()

	if !dataFound {
		return result, fmt.Errorf("required data not found in any script tags")
	}

	return result, nil
}

// extractDataFromScript parses the script content to extract the required fields
func extractDataFromScript(scriptContent string) (types.Result, error) {
	// The content has escaped quotes (\\"), so we need to adjust our regex patterns
	pricingDateRegex := regexp.MustCompile(`\\\"pricingDataDate\\\":\\\"([^\\]+)\\\"`)
	totalAssetRegex := regexp.MustCompile(`\\\"totalAssetInTrust\\\":\\\"([^\\]+)\\\"`)

	pricingDateMatch := pricingDateRegex.FindStringSubmatch(scriptContent)
	totalAssetMatch := totalAssetRegex.FindStringSubmatch(scriptContent)

	if len(pricingDateMatch) < 2 || len(totalAssetMatch) < 2 {
		return types.Result{}, fmt.Errorf("required fields not found in script content")
	}

	// Extract the values
	pricingDateStr := pricingDateMatch[1]
	totalAssetStr := totalAssetMatch[1]

	// Parse the total asset value
	totalAssetStr = strings.ReplaceAll(totalAssetStr, "$", "")
	totalAssetStr = strings.ReplaceAll(totalAssetStr, ",", "")

	totalAsset, err := strconv.ParseFloat(totalAssetStr, 64)
	if err != nil {
		return types.Result{}, fmt.Errorf("failed to parse totalAsset '%s': %w", totalAssetStr, err)
	}

	// Parse the date
	layout := "01/02/2006"
	parsedTime, err := time.Parse(layout, pricingDateStr)
	if err != nil {
		return types.Result{}, fmt.Errorf("failed to parse date '%s': %w", pricingDateStr, err)
	}

	return types.Result{
		TotalAsset: totalAsset,
		Date:       parsedTime,
	}, nil
}
