package cache

import "github.com/jmiguelrc/sharetracker/api"

var marketTickerPriceCache = make(map[string]api.MarketPrice)

// Returns the current market price, caching the request
func GetCurrentMarketPrice(ticker string) api.MarketPrice {
	if val, ok := marketTickerPriceCache[ticker]; ok {
		return val
	}
	val := api.GetCurrentMarketPrice(ticker)
	marketTickerPriceCache[ticker] = val
	return val
}
