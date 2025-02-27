package main

var (
	tickerDetails = map[string]tickerDetail{
		"ARKB": {Asset: "BTC", Description: "Ark 21Shares", Note: "ARKB holdings are updated 10+ hours after the close of trading"}, // ARK 21Shares Bitcoin ETF
		"BITB": {Asset: "BTC", Description: "Bitwise", Note: "BITB holdings are updated 4.5+ hours after the close of trading"},     // Bitwise Bitcoin ETF
		"BRRR": {Asset: "BTC", Description: "Valkyrie", Note: "BRRR holdings are updated 10+ hours after the close of trading"},     // Valkyrie Bitcoin Fund
		"BTC":  {Asset: "BTC", Description: "Grayscale (Mini)", Note: "BTC holdings are updated 1 day late", Delayed: true},         // Grayscale Bitcoin Mini Trust
		"BTCW": {Asset: "BTC", Description: "WisdomTree", Note: ""},                                                                 // WisdomTree Bitcoin Fund
		"DEFI": {Asset: "BTC", Description: "Hashdex", Note: ""},                                                                    // Hashdex Bitcoin ETF
		"EZBC": {Asset: "BTC", Description: "Franklin", Note: "EZBC holdings are updated 5.5+ hours after the close of trading"},    // Franklin Bitcoin ETF
		"FBTC": {Asset: "BTC", Description: "Fidelity", Note: "FBTC holdings are updated 16+ hours after the close of trading"},     // Fidelity Wise Origin Bitcoin Fund
		"GBTC": {Asset: "BTC", Description: "Grayscale", Note: "GBTC holdings are updated 1 day late", Delayed: true},               // Grayscale Bitcoin Trust
		"HODL": {Asset: "BTC", Description: "VanEck", Note: "HODL holdings are updated 1 day late", Delayed: true},                  // VanEck Bitcoin Trust
		"IBIT": {Asset: "BTC", Description: "BlackRock", Note: "IBIT holdings are updated 13+ hours after the close of trading"},    // iShares Bitcoin Trust
		"CETH": {Asset: "ETH", Description: "21Shares", Note: "CETH holdings are updated 10+ hours after the close of trading"},     // 21Shares Core Ethereum ETF
		"ETH":  {Asset: "ETH", Description: "Grayscale (Mini)", Note: "ETH holdings are updated 1 day late", Delayed: true},         // Grayscale Ethereum Mini Trust
		"ETHA": {Asset: "ETH", Description: "BlackRock", Note: "ETHA holdings are updated 13+ hours after the close of trading"},    // BlackRock iShares Ethereum Trust ETF
		"ETHE": {Asset: "ETH", Description: "Grayscale", Note: "ETHE holdings are updated 1 day late", Delayed: true},               // Grayscale Ethereum Trust
		"ETHV": {Asset: "ETH", Description: "VanEck", Note: "ETHV holdings are updated 1 day late", Delayed: true},                  // VanEck Ethereum ETF
		"ETHW": {Asset: "ETH", Description: "Bitwise", Note: "ETHW holdings are updated 4.5+ hours after the close of trading"},     // Bitwise Ethereum ETF
		"EZET": {Asset: "ETH", Description: "Franklin", Note: "EZET holdings are updated 5.5+ hours after the close of trading"},    // Franklin Ethereum ETF
		"FETH": {Asset: "ETH", Description: "Fidelity", Note: "FETH holdings are updated 16+ hours after the close of trading"},     // Fidelity Ethereum Fund
	}

	assetDetails = map[string]assetDetail{
		"BTC": {Units: "BTC", UnitsLong: "Bitcoin", MinAssetDiff: 1.0},
		"ETH": {Units: "ETH", UnitsLong: "Ether", MinAssetDiff: 10.0},
	}

	nonTradingHolidays = map[string]bool{
		"2025-04-18": true,
		"2025-05-26": true,
		"2025-06-19": true,
		"2025-07-04": true,
		"2025-09-01": true,
		"2025-11-27": true,
		"2025-12-25": true,
		"2026-01-01": true,
		"2026-01-19": true,
		"2026-02-16": true,
		"2026-04-03": true,
		"2026-05-25": true,
		"2026-06-19": true,
		"2026-07-03": true,
		"2026-09-07": true,
		"2026-11-26": true,
		"2026-12-25": true,
		"2027-01-01": true,
		"2027-01-18": true,
		"2027-02-15": true,
		"2027-03-26": true,
		"2027-05-31": true,
		"2027-06-18": true,
		"2027-07-05": true,
		"2027-09-06": true,
		"2027-11-25": true,
		"2027-12-24": true,
		"2028-01-17": true,
		"2028-02-21": true,
		"2028-04-14": true,
		"2028-05-29": true,
		"2028-06-19": true,
		"2028-07-04": true,
		"2028-09-04": true,
		"2028-11-23": true,
		"2028-12-25": true,
	}
)
