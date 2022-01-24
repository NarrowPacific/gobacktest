package gobacktest

import (
	"math"
)

type SettledTrade struct {
	Orders                  []FillEvent
	Qty                     int64
	Profit                  float64
	ProfitPercent           float64
	CumulativeProfit        float64
	CumulativeProfitPercent float64
	RunUp                   float64
	RunUpPercent            float64
	DrawDown                float64
	DrawDownPercent         float64
}

func (s *SettledTrade) UpdateCumulativeValues(settledTrades []SettledTrade, initialCash float64) {
	if len(settledTrades) > 0 {
		lastSettledTrade := settledTrades[len(settledTrades)-1]
		s.CumulativeProfit = lastSettledTrade.CumulativeProfit + s.Profit
		s.CumulativeProfit = math.Round(s.CumulativeProfit*math.Pow10(DP)) / math.Pow10(DP)
		cumulativeProfitPercent := s.Profit / (lastSettledTrade.CumulativeProfit + initialCash) * 100
		s.CumulativeProfitPercent = math.Round(cumulativeProfitPercent*math.Pow10(DP)) / math.Pow10(DP)
	} else {
		s.CumulativeProfit = s.Profit
		cumulativeProfitPercent := s.Profit / (s.Profit + initialCash) * 100
		s.CumulativeProfitPercent = math.Round(cumulativeProfitPercent*math.Pow10(DP)) / math.Pow10(DP)
	}
}
