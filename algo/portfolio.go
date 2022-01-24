package algo

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

type investedAlgo struct {
	gbt.Algo
	symbols []string
}

// IsInvested check if the portfolio hols a position for the given symbol.
func IsInvested(symbols ...string) gbt.AlgoHandler {
	return &investedAlgo{symbols: symbols}
}

// Run runs the algo, returns the bool value of the algo
func (algo investedAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	portfolio, _ := s.Portfolio()
	event, _ := s.Event()

	// if no specified symbol use symbol of current event
	if len(algo.symbols) == 0 {
		symbol := event.Symbol()
		if _, ok := portfolio.IsInvested(symbol); !ok {
			return false, nil
		}

		return true, nil
	}

	// symbols specified
	var invested bool
	for _, symbol := range algo.symbols {
		if _, ok := portfolio.IsInvested(symbol); ok {
			invested = true
		}
	}

	return invested, nil
}

type notInvestedAlgo struct {
	gbt.Algo
	symbols []string
}

// NotInvested check if the portfolio holds no position for the given symbol.
func NotInvested(symbols ...string) gbt.AlgoHandler {
	return &notInvestedAlgo{symbols: symbols}
}

// Run runs the algo, returns the bool value of the algo
func (algo notInvestedAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	portfolio, _ := s.Portfolio()
	event, _ := s.Event()

	// if no specified symbol use symbol of current event
	if len(algo.symbols) == 0 {
		symbol := event.Symbol()
		if _, ok := portfolio.IsInvested(symbol); ok {
			return false, nil
		}

		return true, nil
	}

	// symbols specified
	var invested bool
	for _, symbol := range algo.symbols {
		if _, ok := portfolio.IsInvested(symbol); ok {
			invested = true
		}
	}

	return !invested, nil
}

type longAlgo struct {
	gbt.Algo
	symbols []string
}

// IsInvested check if the portfolio hols a position for the given symbol.
func IsLong(symbols ...string) gbt.AlgoHandler {
	return &longAlgo{symbols: symbols}
}

// Run runs the algo, returns the bool value of the algo
func (algo longAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	portfolio, _ := s.Portfolio()
	event, _ := s.Event()

	// if no specified symbol use symbol of current event
	if len(algo.symbols) == 0 {
		symbol := event.Symbol()
		if _, ok := portfolio.IsLong(symbol); !ok {
			return false, nil
		}

		return true, nil
	}

	// symbols specified
	var isLong bool
	for _, symbol := range algo.symbols {
		if _, ok := portfolio.IsLong(symbol); ok {
			isLong = true
		}
	}

	return isLong, nil
}

type shortAlgo struct {
	gbt.Algo
	symbols []string
}

// IsInvested check if the portfolio hols a position for the given symbol.
func IsShort(symbols ...string) gbt.AlgoHandler {
	return &shortAlgo{symbols: symbols}
}

// Run runs the algo, returns the bool value of the algo
func (algo shortAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	portfolio, _ := s.Portfolio()
	event, _ := s.Event()

	// if no specified symbol use symbol of current event
	if len(algo.symbols) == 0 {
		symbol := event.Symbol()
		if _, ok := portfolio.IsShort(symbol); !ok {
			return false, nil
		}

		return true, nil
	}

	// symbols specified
	var isShort bool
	for _, symbol := range algo.symbols {
		if _, ok := portfolio.IsShort(symbol); ok {
			isShort = true
		}
	}

	return isShort, nil
}
