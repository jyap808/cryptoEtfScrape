package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jyap808/cryptoEtfScrape/funds"
	"github.com/jyap808/cryptoEtfScrape/referencerates"
	"github.com/jyap808/cryptoEtfScrape/referencerates/cmeny"
	"github.com/jyap808/cryptoEtfScrape/types"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	gotwiTypes "github.com/michimani/gotwi/tweet/managetweet/types"
)

type payload struct {
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Embeds    []embed `json:"embeds"`
}

type embed struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type manualData struct {
	Ticker string
	Result types.Result
}

// Hold details for each ticker
type tickerDetail struct {
	Asset       string
	Description string
	Note        string
	Delayed     bool
}

type assetDetail struct {
	Units        string
	UnitsLong    string
	MinAssetDiff float64 // Skip X post when the difference is under this threshold
}

var (
	webhookURL string

	avatarUsername string
	avatarURL      string

	listenPort int

	// track
	tickerResults         = map[string]types.Result{}
	tickerResultsOverride = map[string]types.Result{}
	asset_rrs             cmeny.ReferenceRates

	// polling intervals
	pollMinutes  int = 5
	backoffHours int = 12
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

func main() {
	flag.StringVar(&webhookURL, "webhookURL", "https://discord.com/api/webhooks/", "Webhook URL")
	flag.StringVar(&avatarUsername, "avatarUsername", "Annalee Call", "Avatar username")
	flag.StringVar(&avatarURL, "avatarURL", "https://static1.personality-database.com/profile_images/6604632de9954b4d99575e56404bd8b7.png", "Avatar image URL")
	flag.IntVar(&listenPort, "listenPort", 8081, "Listen port")
	flag.Parse()

	// Initialize empty tickerResult
	var wg sync.WaitGroup
	wgCount := 0
	for ticker := range tickerDetails {
		tickerResults[ticker] = types.Result{}
		tickerResultsOverride[ticker] = types.Result{}
		wgCount++
	}

	// Reference Rates handler
	go handleReferenceRates()

	// Increment the WaitGroup counter for each scraping function
	wg.Add(wgCount)

	// Launch goroutines for scraping functions
	for ticker := range tickerDetails {
		go handleFund(&wg, ticker)
	}

	// Manual endpoints
	http.HandleFunc("/override", handleOverride)
	http.HandleFunc("/update", handleUpdate)

	// Start HTTP server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	log.Println("All scraping functions have finished.")
}

// Generic handler
func handleFund(wg *sync.WaitGroup, ticker string) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes

	for {
		var newResult types.Result
		override := false

		// Check if there is a manual override set
		if tickerResultsOverride[ticker].TotalAsset != 0 {
			newResult = tickerResultsOverride[ticker]

			// Clear override
			tickerResultsOverride[ticker] = types.Result{}
			override = true
		} else {
			newResult = funds.Collector(ticker)
		}

		// Check date is valid. Date is optional so we check it is not none
		if !newResult.Date.IsZero() && newResult.Date.Before(tickerResults[ticker].Date) {
			log.Printf("%s new result before current: %+v", ticker, newResult)

			// Backoff for 1 hr or this just will loop
			time.Sleep(time.Hour * time.Duration(1))

			continue
		}

		if newResult.TotalAsset != tickerResults[ticker].TotalAsset && newResult.TotalAsset != 0 {
			if tickerResults[ticker].TotalAsset == 0 {
				// initialize
				tickerResults[ticker] = newResult
				log.Printf("Initialize %s: %+v", ticker, tickerResults[ticker])
			} else {
				// compare
				assetDiff := newResult.TotalAsset - tickerResults[ticker].TotalAsset
				aD := assetDetails[tickerDetails[ticker].Asset]
				assetPrice := asset_rrs.GetReferenceRatePointer(aD.Units)[0].Value
				if tickerDetails[ticker].Delayed {
					assetPrice = asset_rrs.GetReferenceRatePointer(aD.Units)[1].Value
				}
				flowDiff := assetDiff * assetPrice

				header := ticker
				if newResult.Date != (time.Time{}) {
					layout := "01/02/2006"
					formattedTime := newResult.Date.Format(layout)

					header = fmt.Sprintf("%s %s", ticker, formattedTime)
				}

				msg := fmt.Sprintf("%s\nCHANGE %s: %.1f\nTOTAL %s: %.1f\nDETAILS Flow: $%.1f, RR: $%.1f",
					header, aD.UnitsLong, assetDiff, aD.UnitsLong, newResult.TotalAsset,
					flowDiff, asset_rrs.ETHUSD_NY[0].Value)

				postDiscord(msg)

				flowEmoji := "🚀"
				if assetDiff < 0 {
					flowEmoji = "👎"
				}

				note := ""
				if !override {
					note = tickerDetails[ticker].Note
				}

				xMsg := fmt.Sprintf("%s $%s\n\n%s FLOW: %s %s, $%s\n🏦 TOTAL %s in Trust: %s $%s\n\n%s",
					tickerDetails[ticker].Description, ticker,
					flowEmoji, humanize.CommafWithDigits(assetDiff, 2), aD.Units, humanize.CommafWithDigits(flowDiff, 0),
					aD.UnitsLong, humanize.CommafWithDigits(newResult.TotalAsset, 1), aD.Units, note)

				// Reporting threshold check. Get the absolute difference
				absAssetDiff := math.Abs(assetDiff)
				if absAssetDiff > aD.MinAssetDiff {
					postTweet(xMsg)
				}

				tickerResults[ticker] = newResult

				log.Printf("Update %s: %+v", ticker, tickerResults[ticker])

				time.Sleep(time.Hour * time.Duration(backoffHours))
			}
		}

		time.Sleep(time.Minute * time.Duration(pollMinutes))
	}
}

func handleData(w http.ResponseWriter, r *http.Request, updateType string) {
	var data manualData

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	// Set based on updateType
	var results map[string]types.Result
	switch updateType {
	case "override":
		results = tickerResultsOverride
	case "update":
		results = tickerResults
	default:
		log.Println("Invalid update type")
		return
	}

	// Update the corresponding map
	newResult := types.Result{TotalAsset: data.Result.TotalAsset, Date: data.Result.Date}
	results[data.Ticker] = newResult

	// Log and respond with success message
	log.Printf("Data %s %s: %+v", updateType, data.Ticker, results[data.Ticker])
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data %s successful\n", updateType)
}

func handleOverride(w http.ResponseWriter, r *http.Request) {
	handleData(w, r, "override")
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	handleData(w, r, "update")
}

// Initializes and updates daily at 4pm ET (with update buffer)
func handleReferenceRates() {
	// Load the Eastern Time location
	etLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Failed to load Eastern Time location: %v", err)
	}

	for {
		rrs, err := referencerates.CMENYCollect()
		if err != nil {
			log.Println("Get Reference Rates error:", err)
			if len(asset_rrs.BRRNY) == 0 {
				log.Fatalln("Reference Rates failed to initialize. Exiting...")
			}

			// Retry timer
			time.Sleep(time.Hour)
			continue
		}

		log.Printf("Set Reference Rates: %v\n", rrs)
		asset_rrs = rrs

		// Calculate time until next update based on rrs.BRRNY[0].Date
		lastUpdateTime := rrs.BRRNY[0].Date
		now := time.Now().In(etLocation)
		nextUpdateTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 11, 0, 0, etLocation)

		if now.After(nextUpdateTime) {
			nextUpdateTime = nextUpdateTime.Add(24 * time.Hour)
		}

		sleepDuration := time.Until(nextUpdateTime)

		log.Printf("Reference Rates Last update: %v, Next update: %v, Sleeping for: %v\n",
			lastUpdateTime, nextUpdateTime, sleepDuration)
		time.Sleep(sleepDuration)
	}
}

func postDiscord(msg string) {
	blockEmbed := embed{Description: msg}
	embeds := []embed{blockEmbed}
	jsonReq := payload{Username: avatarUsername, AvatarURL: avatarURL, Embeds: embeds}

	jsonStr, _ := json.Marshal(jsonReq)
	log.Println("Discord POST:", string(jsonStr))

	req, _ := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}

func postTweet(msg string) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		log.Println(err)
		return
	}

	p := &gotwiTypes.CreateInput{
		Text: gotwi.String(msg),
	}

	// Replace newline characters with spaces
	logStr := strings.ReplaceAll(msg, "\n", " ")
	log.Println("X Tweet:", logStr)

	_, err = managetweet.Create(context.Background(), c, p)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
