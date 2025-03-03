package franklin

import (
	"bytes"
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
	Portfolio Portfolio `json:"Portfolio"`
}

type Portfolio struct {
	PortfolioData PortfolioData `json:"portfolio"`
}

type PortfolioData struct {
	DailyHoldings []Holding `json:"dailyholdings"`
}

type Holding struct {
	Date     string `json:"asofdate"`
	SECName  string `json:"secname"`
	Quantity string `json:"quantityshrpar"`
}

func CollectWithFundIDAndSearch(fundID int, search string) (result types.Result, err error) {
	// This API key is hard coded on their web site
	url := "https://www.franklintempleton.com/api/pds/price-and-performance?apikey=4ef35821-5244-41bc-a699-0192d002c3d1p&op=Holdings&id=14"

	// JSON payload
	payloadString := fmt.Sprintf(`{"operationName":"Holdings","variables":{"countrycode":"US","languagecode":"en_US","fundid":"%d"},"query":"query Holdings($fundid: String!, $countrycode: String!, $languagecode: String!) {Portfolio(fundid: $fundid countrycode: $countrycode languagecode: $languagecode) {portfolio {dailyholdings {asofdate secname quantityshrpar}}}}"}`, fundID)
	payload := []byte(payloadString)

	// Create a new HTTP client
	client := http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return types.Result{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

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

	// Iterate
	for _, holding := range data.Data.Portfolio.PortfolioData.DailyHoldings {
		if holding.SECName == search {
			// Extract
			totalRaw := holding.Quantity
			inputClean := strings.ReplaceAll(totalRaw, ",", "")
			total, _ := strconv.ParseFloat(inputClean, 64)
			result.TotalAsset = total

			// Define the layout of the input date
			layout := "01/02/2006"
			// Parse the string as a time.Time value
			parsedTime, _ := time.Parse(layout, holding.Date)
			result.Date = parsedTime
			return result, nil
		}
	}

	return types.Result{}, fmt.Errorf("no holding found for search term: %s", search)
}
