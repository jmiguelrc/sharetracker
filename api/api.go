package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// Retuns the current market price a given ticker in Yahoo Finance (https://finance.yahoo.com/)
func GetCurrentMarketPrice(ticker string) float64 {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?region=GB&lang=en-GB&.tsrc=finance", ticker)
	resp, err := resty.New().R().Get(url)
	if err != nil {
		log.Fatalf("Error during HTTP request: %v", err)
	}
	var responseObj = YahooFinanceResponse{}
	err = json.Unmarshal(resp.Body(), &responseObj)
	if err != nil {
		log.Fatalf("Failure parsing json response: %v", err)
	}
	return responseObj.Chart.Result[0].Meta.RegularMarketPrice
}
