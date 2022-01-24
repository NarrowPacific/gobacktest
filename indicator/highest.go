package indicator

import (
	"fmt"

	gbt "github.com/dirkolbrich/gobacktest"
)

// highestAlgo is an algo which calculates the simple moving average.
type highestAlgo struct {
	gbt.Algo
	source   Source
	lookback int
	highest  float64
}

// Highest returns highest value for a given number of lookback
func Highest(source Source, lookback int) gbt.AlgoHandler {
	return &highestAlgo{source: source, lookback: lookback}
}

// Run runs the algo.
func (l *highestAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	data, _ := s.Data()
	event, _ := s.Event()
	symbol := event.Symbol()

	// prepare list of floats
	list := data.List(symbol)
	var bars []*gbt.Bar

	if len(list) < l.lookback {
		return false, fmt.Errorf("invalid value length for indicator highest")
	}

	for i := 0; i < l.lookback; i++ {
		bars = append(bars, list[len(list)-i-1].(*gbt.Bar))
	}

	// calculate Highest
	l.highest = highest(bars, l.source)

	return true, nil
}

// Value returns the value of this Algo.
func (l *highestAlgo) Value() interface{} {
	return l.highest
}

func highest(bars []*gbt.Bar, source Source) float64 {
	highest := getValueBySource(bars[0], source)
	for i := 1; i < len(bars); i++ {
		value := getValueBySource(bars[i], source)
		if value > highest {
			highest = value
		}
	}

	return highest
}
