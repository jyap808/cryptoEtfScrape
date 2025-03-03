package vaneck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jyap808/cryptoEtfScrape/types"
)

type FundData struct {
	Data Data `json:"data"`
}

type Data struct {
	Date string `json:"AsOfDate"`
	Navs []Nav
}

type Nav struct {
	Key   string
	Value string
}

func CollectFromURLAndSearch(url string, search string) (result types.Result, err error) {
	// NOTE: Fix for getting old cached responses from this endpoint
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Result{}, fmt.Errorf("error creating request: %w", err)
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return types.Result{}, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Result{}, fmt.Errorf("error reading response body: %w", err)
	}

	return parseJSON(body, search)
}

func parseJSON(body []byte, search string) (result types.Result, err error) {
	var data FundData
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return types.Result{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Define the layout of the input date
	layout := "01/02/2006"
	// Parse the string as a time.Time value
	parsedTime, _ := time.Parse(layout, data.Data.Date)
	result.Date = parsedTime

	// Iterate
	for _, nav := range data.Data.Navs {
		if nav.Key == search {
			// Extract
			totalRaw := nav.Value
			inputClean := strings.ReplaceAll(totalRaw, ",", "")
			total, _ := strconv.ParseFloat(inputClean, 64)
			result.TotalAsset = total
			return
		}
	}

	return types.Result{}, fmt.Errorf("no holding found for search term: %s", search)
}
