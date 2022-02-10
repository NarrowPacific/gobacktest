package algo

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

type crossoverFixedAlgo struct {
	gbt.Algo
	series, fixed gbt.AlgoHandler
}

func CrossoverFixed(series, fixed gbt.AlgoHandler) gbt.AlgoHandler {
	return &crossoverFixedAlgo{
		series: series,
		fixed:  fixed,
	}
}

func (algo crossoverFixedAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	okFirst, err := algo.series.Run(s)
	if !okFirst || err != nil {
		return false, err
	}

	fixedValue := algo.fixed.Value().(float64)

	result := crossoverFixed(algo.series.Value().([]float64), fixedValue)

	return result, nil
}

func crossoverFixed(series []float64, fixedValue float64) bool {
	if len(series) < 2 {
		return false
	}

	if series[len(series)-1] == 0 || series[len(series)-2] == 0 {
		return false
	}

	return series[len(series)-2] <= fixedValue && series[len(series)-1] > fixedValue
}
