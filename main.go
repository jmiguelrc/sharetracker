package main

import (
	"fmt"
	"log"

	"github.com/jmiguelrc/sharetracker/positioncalc"
	"github.com/leekchan/accounting"
	"github.com/spf13/viper"
)

var gainsTaxRate float64
var positions []positioncalc.Position

func init() {
	viper.SetConfigName(".sharetracker")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}

	viper.SetDefault("gainsTaxRate", 0.28)
	gainsTaxRate = viper.GetFloat64("gainsTaxRate")
	positions = make([]positioncalc.Position, 0)
	err = viper.UnmarshalKey("positions", &positions)
	if err != nil {
		log.Fatalf("Error reading configured positions %s", err)
	}
}

func main() {
	currentStatus := positioncalc.CalcProfitPositions(positions)

	ac := accounting.Accounting{Symbol: "€", Precision: 2}
	fmt.Printf("Total buy value:  %s\n", ac.FormatMoney(currentStatus.TotalBuyValue))
	fmt.Printf("Current Value:    %s\n", ac.FormatMoney(currentStatus.GrossCurrentValue))
	fmt.Printf("Total Net Value:  %s\n", ac.FormatMoney(currentStatus.NetCurrentValue(gainsTaxRate)))
	fmt.Printf("Gross Profit:     %s\n", ac.FormatMoney(currentStatus.GrossProfit()))
	fmt.Printf("Net Profit:       %s\n", ac.FormatMoney(currentStatus.NetProfit(gainsTaxRate)))
	fmt.Printf("YTD Gross Profit: %s\n", ac.FormatMoney(currentStatus.YtdGrossProfit))
}
