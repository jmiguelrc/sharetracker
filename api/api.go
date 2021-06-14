package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type MarketPrice struct {
	CurrentPrice   float64
	YearStartPrice float64
}

// Retuns the current market price a given ticker in Yahoo Finance (https://finance.yahoo.com/)
func GetCurrentMarketPrice(ticker string) MarketPrice {
	response := getTickerInfo(ticker)
	yearStartPrice := response.Chart.Result[0].Indicators.Quote[0].Close[0]
	currentMarketPrice := response.Chart.Result[0].Meta.RegularMarketPrice
	return MarketPrice{CurrentPrice: currentMarketPrice, YearStartPrice: yearStartPrice}
}

func getTickerInfo(ticker string) YahooFinanceResponse {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&interval=1d&useYfid=true&range=ytd", ticker)
	resp, err := resty.New().R().Get(url)
	if err != nil {
		log.Fatalf("Error during HTTP request: %v", err)
	}
	if resp.RawResponse.StatusCode != 200 {
		log.Fatalf("Error response: %d; %v", resp.RawResponse.StatusCode, resp.RawResponse)
	}
	var responseObj = YahooFinanceResponse{}
	err = json.Unmarshal(resp.Body(), &responseObj)
	if err != nil {
		log.Fatalf("Failure parsing json response: %v", err)
	}
	return responseObj
}
