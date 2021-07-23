package positioncalc

import (
	"time"

	"github.com/jmiguelrc/sharetracker/cache"
)

type Position struct {
	Ticker    string
	NumShares float64
	BuyPrice  float64
	Date      string
}

type CurrentProfit struct {
	TotalBuyValue     float64
	GrossCurrentValue float64
	NetCurrentValue   float64
	NetProfit         float64
	YtdGrossProfit    float64
}

func CalcPositionResult(pos Position, gainsTaxRate float64) CurrentProfit {
	var currentStatus CurrentProfit

	buyDate, _ := time.Parse("02-01-2006", pos.Date)
	wasBoughThisYear := buyDate.Year() == time.Now().Year()

	currentStatus.TotalBuyValue = pos.NumShares * pos.BuyPrice
	marketPrice := cache.GetCurrentMarketPrice(pos.Ticker)
	currentStatus.GrossCurrentValue = pos.NumShares * marketPrice.CurrentPrice

	ytdPriceReference := marketPrice.YearStartPrice
	if wasBoughThisYear {
		ytdPriceReference = pos.BuyPrice
	}
	currentStatus.YtdGrossProfit = pos.NumShares * (marketPrice.CurrentPrice - ytdPriceReference)

	return currentStatus
}

func CalcProfitPositions(positions []Position, gainsTaxRate float64) CurrentProfit {
	var currentStatus CurrentProfit
	for _, position := range positions {
		positionCurStatus := CalcPositionResult(position, gainsTaxRate)
		currentStatus.GrossCurrentValue += positionCurStatus.GrossCurrentValue
		currentStatus.TotalBuyValue += positionCurStatus.TotalBuyValue
		currentStatus.YtdGrossProfit += positionCurStatus.YtdGrossProfit
	}

	currentStatus.NetProfit = (1 - gainsTaxRate) * (currentStatus.GrossCurrentValue - currentStatus.TotalBuyValue)
	currentStatus.NetCurrentValue = currentStatus.TotalBuyValue + currentStatus.NetProfit

	return currentStatus
}
