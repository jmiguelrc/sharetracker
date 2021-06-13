package cache

import "github.com/jmiguelrc/sharetracker/api"

var marketTickerPriceCache = make(map[string]float64)

// Returns the current market price, caching the request
func GetCurrentMarketPrice(ticker string) float64 {
	if val, ok := marketTickerPriceCache[ticker]; ok {
		return val
	}
	val := api.GetCurrentMarketPrice(ticker)
	marketTickerPriceCache[ticker] = val
	return val
}
