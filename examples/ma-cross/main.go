package main

import (
	"fmt"

	gbt "github.com/dirkolbrich/gobacktest"
	"github.com/dirkolbrich/gobacktest/algo"
	"github.com/dirkolbrich/gobacktest/data"
	"github.com/dirkolbrich/gobacktest/indicator"
)

func main() {
	// initiate new backtester
	test := gbt.New()
	// set the portfolio with initial cash and a default size and risk manager
	portfolio := &gbt.Portfolio{}
	portfolio.SetInitialCash(100000)

	sizeManager := &gbt.Size{DefaultSize: 1, DefaultValue: 1000000000}
	portfolio.SetSizeManager(sizeManager)

	riskManager := &gbt.Risk{}
	portfolio.SetRiskManager(riskManager)

	test.SetPortfolio(portfolio)

	// define and load symbols
	var symbol = "BTCUSDT"
	var symbols = []string{symbol}
	test.SetSymbols(symbols)

	// create data provider and load data into the backtest
	data := &data.BarEventFromCSVFile{FileDir: "examples/testdata/bar/"}
	data.Load(symbols)
	test.SetData(data)

	// create a new strategy with an algo stack and load into the backtest
	strategy := gbt.NewStrategy("ema-cross")
	fast := 12
	slow := 26
	strategy.SetAlgo(
		algo.If(
			// condition
			algo.And(
				algo.Crossover(indicator.EMA(fast), indicator.EMA(slow)),
				algo.NotInvested(),
			),
			// action
			algo.CreateSignal("buy"), // create a buy signal
		),
		algo.If(
			// condition
			algo.And(
				algo.Crossunder(indicator.EMA(fast), indicator.EMA(slow)),
				algo.NotInvested(),
			),
			// action
			algo.CreateSignal("sell"), // create a sell signal
		),
		algo.If(
			// condition exit long
			algo.And(
				algo.Crossunder(indicator.EMA(fast), indicator.EMA(slow)),
				algo.IsLong(),
			),
			// action
			algo.CreateSignal("exit"), // create a exit signal
		),
		algo.If(
			// condition exit short
			algo.And(
				algo.Crossover(indicator.EMA(fast), indicator.EMA(slow)),
				algo.IsShort(),
			),
			// action
			algo.CreateSignal("exit"), // create a exit signal
		),
	)

	// create an asset and append to strategy
	strategy.SetChildren(gbt.NewAsset(symbol))

	// load the strategy into the backtest
	test.SetStrategy(strategy)

	// run the backtest
	err := test.Run()
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// print the result of the test
	test.Stats().PrintResult()
}
