package main

import (
	"fmt"
	"log"

	"github.com/jmiguelrc/sharetracker/cache"
	"github.com/leekchan/accounting"
	"github.com/spf13/viper"
)

type Position struct {
	Ticker    string
	NumShares float64
	BuyPrice  float64
}

var gainsTaxRate float64
var positions []Position

func init() {
	viper.SetConfigName("portfolio")
	viper.AddConfigPath("$HOME/.sharetracker")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}

	viper.SetDefault("gainsTaxRate", 0.28)
	gainsTaxRate = viper.GetFloat64("gainsTaxRate")
	positions = make([]Position, 0)
	viper.UnmarshalKey("positions", &positions)
}

func main() {
	investedValue := 0.0
	currentMarketValue := 0.0

	for _, position := range positions {
		investedValue += position.NumShares * position.BuyPrice
		currentMarketValue += position.NumShares * cache.GetCurrentMarketPrice(position.Ticker)
	}

	netProfit := (1 - gainsTaxRate) * (currentMarketValue - investedValue)
	totalNetValue := investedValue + netProfit

	ac := accounting.Accounting{Symbol: "€", Precision: 2}
	fmt.Printf("Total buy value: %s\n", ac.FormatMoney(investedValue))
	fmt.Printf("Current Value:   %s\n", ac.FormatMoney(currentMarketValue))
	fmt.Printf("Gross Profit:    %s\n", ac.FormatMoney(currentMarketValue-investedValue))
	fmt.Printf("Net Profit:      %s\n", ac.FormatMoney(netProfit))
	fmt.Printf("Total Net Value: %s\n", ac.FormatMoney(totalNetValue))
}
