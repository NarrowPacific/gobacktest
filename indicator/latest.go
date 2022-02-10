package indicator

import (
	gbt "github.com/dirkolbrich/gobacktest"
)

// latestAlgo is an algo which return latest price
type latestAlgo struct {
	gbt.Algo
	source Source
	latest float64
}

// Latest returns latest value
func Latest(source Source) gbt.AlgoHandler {
	return &latestAlgo{source: source}
}

// Run runs the algo.
func (l *latestAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	data, _ := s.Data()
	event, _ := s.Event()
	symbol := event.Symbol()
	l.latest = getValueBySource(data.Latest(symbol).(*gbt.Bar), l.source)

	return true, nil
}

// Value returns the value of this Algo.
func (l *latestAlgo) Value() interface{} {
	return l.latest
}
