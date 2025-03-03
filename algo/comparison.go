package algo

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

type biggerThanAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

// BiggerThan compares the value of the two containing algos.
func BiggerThan(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &biggerThanAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo biggerThanAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	var firstValue, secondValue float64
	if firstValues, ok := algo.first.Value().([]float64); ok {
		firstValue = firstValues[len(firstValues)-1]
	} else {
		firstValue = algo.first.Value().(float64)
	}
	if secondValues, ok := algo.second.Value().([]float64); ok {
		secondValue = secondValues[len(secondValues)-1]
	} else {
		secondValue = algo.second.Value().(float64)
	}

	result := firstValue > secondValue

	return result, nil
}

type biggerThanOrEqualToAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

// BiggerThanOrEqualTo compares the value of the two containing algos.
func BiggerThanOrEqualTo(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &biggerThanOrEqualToAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo biggerThanOrEqualToAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	var firstValue, secondValue float64
	if firstValues, ok := algo.first.Value().([]float64); ok {
		firstValue = firstValues[len(firstValues)-1]
	} else {
		firstValue = algo.first.Value().(float64)
	}
	if secondValues, ok := algo.second.Value().([]float64); ok {
		secondValue = secondValues[len(secondValues)-1]
	} else {
		secondValue = algo.second.Value().(float64)
	}

	result := firstValue >= secondValue

	return result, nil
}

type smallerThanAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

// SmallerThan compares if the value of the first algo is smaller than second.
func SmallerThan(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &smallerThanAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo smallerThanAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	var firstValue, secondValue float64
	if firstValues, ok := algo.first.Value().([]float64); ok {
		firstValue = firstValues[len(firstValues)-1]
	} else {
		firstValue = algo.first.Value().(float64)
	}
	if secondValues, ok := algo.second.Value().([]float64); ok {
		secondValue = secondValues[len(secondValues)-1]
	} else {
		secondValue = algo.second.Value().(float64)
	}

	result := firstValue < secondValue

	return result, nil
}

type smallerThanOrEqualToAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

// SmallerThanOrEqualTo compares if the value of the first algo is smaller than or equal to second.
func SmallerThanOrEqualTo(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &smallerThanOrEqualToAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo smallerThanOrEqualToAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	var firstValue, secondValue float64
	if firstValues, ok := algo.first.Value().([]float64); ok {
		firstValue = firstValues[len(firstValues)-1]
	} else {
		firstValue = algo.first.Value().(float64)
	}
	if secondValues, ok := algo.second.Value().([]float64); ok {
		secondValue = secondValues[len(secondValues)-1]
	} else {
		secondValue = algo.second.Value().(float64)
	}

	result := firstValue <= secondValue

	return result, nil
}

type equalAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
}

// Equal compares the value of two algos.
func Equal(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &equalAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo equalAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	var firstValue, secondValue float64
	if firstValues, ok := algo.first.Value().([]float64); ok {
		firstValue = firstValues[len(firstValues)-1]
	} else {
		firstValue = algo.first.Value().(float64)
	}
	if secondValues, ok := algo.second.Value().([]float64); ok {
		secondValue = secondValues[len(secondValues)-1]
	} else {
		secondValue = algo.second.Value().(float64)
	}

	result := firstValue == secondValue

	return result, nil
}
