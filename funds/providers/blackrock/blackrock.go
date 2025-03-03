package blackrock

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jyap808/cryptoEtfScrape/types"
)

type FundData struct {
	AaData [][]interface{} `json:"aaData"`
}

type DisplayData struct {
	Display string  `json:"display"`
	Raw     float64 `json:"raw"`
}

func CollectFromURLAndTicker(url string, ticker string) (result types.Result, err error) {
	// Create a new HTTP client
	client := http.Client{}

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

	// Trim any leading characters that may cause the issue
	bodyStr := string(body)
	bodyStr = strings.TrimLeftFunc(bodyStr, func(r rune) bool {
		return r != '{' && r != '['
	})

	return parseJSON(bodyStr, ticker)
}

func parseJSON(bodyStr string, ticker string) (result types.Result, err error) {
	var data FundData
	if err := json.Unmarshal([]byte(bodyStr), &data); err != nil {
		return types.Result{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Iterate through the funds and find the one with ticker
	for _, fund := range data.AaData {
		if len(fund) > 6 && fund[0] == ticker {
			sharesJSON, err := json.Marshal(fund[6])
			if err != nil {
				return types.Result{}, fmt.Errorf("error marshalling shares data: %v", err)
			}

			var shares DisplayData
			if err := json.Unmarshal(sharesJSON, &shares); err != nil {
				return types.Result{}, fmt.Errorf("error unmarshalling shares data: %v", err)
			}

			result.TotalAsset = shares.Raw
			return result, nil
		}
	}

	return result, nil
}
