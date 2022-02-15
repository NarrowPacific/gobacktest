package indicator

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

type multiplyAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
	value         float64
}

func Multiply(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &multiplyAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo *multiplyAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	algo.value = firstValue * secondValue

	return true, nil
}

// Value returns the value of this Algo.
func (algo *multiplyAlgo) Value() interface{} {
	return algo.value
}

type divideAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
	value         float64
}

func Divide(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &divideAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo *divideAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	algo.value = firstValue / secondValue

	return true, nil
}

// Value returns the value of this Algo.
func (algo *divideAlgo) Value() interface{} {
	return algo.value
}

type addAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
	value         float64
}

func Add(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &addAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo *addAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	algo.value = firstValue + secondValue

	return true, nil
}

// Value returns the value of this Algo.
func (algo *addAlgo) Value() interface{} {
	return algo.value
}

type subtractAlgo struct {
	gbt.Algo
	first, second gbt.AlgoHandler
	value         float64
}

func Subtract(first, second gbt.AlgoHandler) gbt.AlgoHandler {
	return &subtractAlgo{
		first:  first,
		second: second,
	}
}

// Run runs the algo, returns the bool value of the algo
func (algo *subtractAlgo) Run(s gbt.StrategyHandler) (bool, error) {
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

	algo.value = firstValue - secondValue

	return true, nil
}

// Value returns the value of this Algo.
func (algo *subtractAlgo) Value() interface{} {
	return algo.value
}
