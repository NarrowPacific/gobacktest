package indicator

import (
	"fmt"

	gbt "github.com/dirkolbrich/gobacktest"
)

type Source int

const (
	OPEN Source = iota // 0
	HIGH
	LOW
	CLOSE
)

// lowestAlgo is an algo which calculates the simple moving average.
type lowestAlgo struct {
	gbt.Algo
	source   Source
	lookback int
	lowest   float64
}

// Lowest returns lowest value for a given number of lookback
func Lowest(source Source, lookback int) gbt.AlgoHandler {
	return &lowestAlgo{source: source, lookback: lookback}
}

// Run runs the algo.
func (l *lowestAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	data, _ := s.Data()
	event, _ := s.Event()
	symbol := event.Symbol()

	// prepare list of floats
	list := data.List(symbol)
	var bars []*gbt.Bar

	if len(list) < l.lookback {
		return false, fmt.Errorf("invalid value length for indicator lowest")
	}

	for i := 0; i < l.lookback; i++ {
		bars = append(bars, list[len(list)-i-1].(*gbt.Bar))
	}

	// calculate Lowest
	l.lowest = lowest(bars, l.source)

	return true, nil
}

// Value returns the value of this Algo.
func (l *lowestAlgo) Value() interface{} {
	return l.lowest
}

func lowest(bars []*gbt.Bar, source Source) float64 {
	lowest := getValueBySource(bars[0], source)
	for i := 1; i < len(bars); i++ {
		value := getValueBySource(bars[i], source)
		if value < lowest {
			lowest = value
		}
	}

	return lowest
}

func getValueBySource(bar *gbt.Bar, source Source) float64 {
	var value float64
	switch source {
	case OPEN:
		value = bar.Open
	case HIGH:
		value = bar.High
	case LOW:
		value = bar.Low
	case CLOSE:
		value = bar.Close
	}

	return value
}
