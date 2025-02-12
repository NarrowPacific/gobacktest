package algo

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

type crossoverAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

func Crossover(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &crossoverAlgo{
		first:  first,
		second: second,
	}
}

func (algo crossoverAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	okFirst, err := algo.first.Run(s)
	if err != nil {
		return false, err
	}

	okSecond, err := algo.second.Run(s)
	if err != nil {
		return false, err
	}

	if !okFirst || !okSecond {
		return false, nil
	}

	result := crossover(algo.first.Value().([]float64), algo.second.Value().([]float64))

	return result, nil
}

func crossover(series1, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	if series1[len(series1)-1] == 0 || series1[len(series1)-2] == 0 || series2[len(series2)-1] == 0 || series2[len(series2)-2] == 0 {
		return false
	}

	return series1[len(series1)-2] <= series2[len(series2)-2] && series1[len(series1)-1] > series2[len(series2)-1]
}
