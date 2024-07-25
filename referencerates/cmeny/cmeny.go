/*
The CME CF Reference Rates – New York Variant is a once a day (4pm ET)
benchmark prices measured in US dollars per asset.

https://www.cmegroup.com/markets/cryptocurrencies/cme-cf-cryptocurrency-benchmarks.html
*/
package cmeny

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ReferenceRate struct {
	Value float64   `json:",string"`
	Date  time.Time `json:"date"`
}

type ReferenceRates struct {
	BRRNY     [5]ReferenceRate `json:"BRRNY"`
	ETHUSD_NY [5]ReferenceRate `json:"ETHUSD_NY"`
}

// GetReferenceRatePointer returns a pointer to the specified reference rate array
func (rr *ReferenceRates) GetReferenceRatePointer(asset string) *[5]ReferenceRate {
	switch asset {
	case "BTC":
		return &rr.BRRNY
	case "ETH":
		return &rr.ETHUSD_NY
	default:
		return nil
	}
}

// Custom unmarshalling function for time.Time field
func (rr *ReferenceRate) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Value float64 `json:",string"`
		Date  string  `json:"date"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", tmp.Date, time.UTC)
	if err != nil {
		return err
	}
	rr.Value = tmp.Value
	rr.Date = date
	return nil
}

func GetReferenceRates() (referenceRates ReferenceRates, err error) {
	url := "https://www.cmegroup.com/services/cryptocurrencies/reference-rates"

	// Create a new HTTP client
	client := http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return ReferenceRates{}, err
	}

	// Set headers
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15")

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error performing request:", err)
		return ReferenceRates{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ReferenceRates{}, err
	}

	// Parse JSON data into struct
	var data map[string]ReferenceRates
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return ReferenceRates{}, err
	}

	return data["referenceRates"], nil
}
