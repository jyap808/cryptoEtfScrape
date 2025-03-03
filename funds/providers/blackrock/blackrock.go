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

type Fund struct {
	Ticker string
	Shares types.Result
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
		if len(fund) > 0 && fund[0] == ticker {
			// Extract the "Shares" field
			sharesMap, ok := fund[6].(map[string]interface{})
			if ok {
				sharesRaw, _ := sharesMap["raw"].(float64)
				result.TotalAsset = sharesRaw
				return result, nil
			}
		}
	}

	return result, nil
}
