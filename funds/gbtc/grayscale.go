package gbtc

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/jyap808/cryptoEtfScrape/types"
)

type nextData struct {
	Props struct {
		PageProps struct {
			Page struct {
				Includes map[string]interface{}
			}
		}
	}
}

func Collect() (result types.Result) {
	// creating a new Colly instance
	c := colly.NewCollector()

	// Set up a callback to be executed when the HTML body is found
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Get the content of the __NEXT_DATA__ script tag
		nextDataContent := e.DOM.Find("#__NEXT_DATA__").Text()

		// Parse the content as JSON
		var data nextData
		err := json.NewDecoder(strings.NewReader(nextDataContent)).Decode(&data)
		if err != nil {
			log.Println(err)
		}

		// Access the "includes" field
		includesData := data.Props.PageProps.Page.Includes

		// Search for the value within "includes"
		result, err = findResultsInIncludes(includesData)
		if err != nil {
			log.Println(err)
		}
	})

	// visiting the target page
	c.Visit("https://etfs.grayscale.com/gbtc")

	c.Wait()

	return result
}

// findResultsInIncludes searches for the unique field within "includes"
func findResultsInIncludes(includesData map[string]interface{}) (types.Result, error) {
	result := &types.Result{}

	for _, value := range includesData {
		// Assuming the value is a map[string]interface{}
		include, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		// Search for "totalAssetInTrustRaw" within each include
		totalAssetInTrustRaw, found := include["totalAssetInTrust"].(string)
		if found {
			inputClean := strings.ReplaceAll(totalAssetInTrustRaw, ",", "")
			totalAssetInTrust, _ := strconv.ParseFloat(inputClean, 64)

			// Define the layout of the input date
			layout := "01/02/2006"
			// Parse the string as a time.Time value
			parsedTime, _ := time.Parse(layout, include["date"].(string))

			result.TotalAsset = totalAssetInTrust
			result.Date = parsedTime

			return *result, nil
		}
	}

	return types.Result{}, fmt.Errorf("totalAssetInTrust not found within 'includes'")
}
