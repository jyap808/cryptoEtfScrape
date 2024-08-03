package twentyoneshares

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jyap808/cryptoEtfScrape/types"
)

func CollectFromURL(url string) (result types.Result) {
	// Create a new HTTP client
	client := http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh)")

	// Perform the request
	resp, err := client.Do(req)
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

	r := csv.NewReader(strings.NewReader(string(body)))

	for i := 0; i < 2; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		// CSV record validity check
		if len(record) < 6 {
			log.Printf("CETH: Invalid record length: expected at least 6 fields, got %d", len(record))
			return
		}

		if i == 1 {
			dateRaw := record[1]
			// Define the layout of the input date
			layout := "01/02/2006"
			// Parse the string as a time.Time value
			parsedTime, _ := time.Parse(layout, dateRaw)

			total, _ := strconv.ParseFloat(record[4], 64)

			result = types.Result{Date: parsedTime, TotalAsset: total}
			return
		}
	}

	return result
}
