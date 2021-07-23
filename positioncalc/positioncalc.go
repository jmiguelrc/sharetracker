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
	YtdGrossProfit    float64
}

func (p *CurrentProfit) NetCurrentValue(gainsTaxRate float64) float64 {
	return p.TotalBuyValue + p.NetProfit(gainsTaxRate)
}

func (p *CurrentProfit) NetProfit(gainsTaxRate float64) float64 {
	if p.GrossProfit() > 0 {
		return p.GrossProfit() * (1 - gainsTaxRate)
	} else {
		return p.GrossProfit()
	}
}

func (p *CurrentProfit) GrossProfit() float64 {
	return p.GrossCurrentValue - p.TotalBuyValue
}

func CalcPositionResult(pos Position, gainsTaxRate float64) CurrentProfit {
	var currentStatus CurrentProfit

	currentStatus.TotalBuyValue = pos.NumShares * pos.BuyPrice
	marketPrice := cache.GetCurrentMarketPrice(pos.Ticker)
	currentStatus.GrossCurrentValue = pos.NumShares * marketPrice.CurrentPrice

	currentStatus.YtdGrossProfit = pos.NumShares * (marketPrice.CurrentPrice - getYtdBasePrice(pos, marketPrice.YearStartPrice))

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

	return currentStatus
}

func getYtdBasePrice(pos Position, yearStartPrice float64) float64 {
	buyDate, err := time.Parse("02-01-2006", pos.Date)

	wasBoughThisYear := err == nil && buyDate.Year() == time.Now().Year()

	if wasBoughThisYear {
		return pos.BuyPrice
	}
	return yearStartPrice
}
