package indicator

import (
	"fmt"

	gbt "github.com/dirkolbrich/gobacktest"
	"github.com/dirkolbrich/gobacktest/ta"
)

// emaAlgo is an algo which calculates the simple moving average.
type emaAlgo struct {
	gbt.Algo
	period int
	ema    float64
	values []float64
}

// EMA returns a ema algo ready to use.
func EMA(i int) gbt.AlgoHandler {
	return &emaAlgo{period: i}
}

// Run runs the algo.
func (a *emaAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	data, _ := s.Data()
	event, _ := s.Event()
	symbol := event.Symbol()

	// prepare list of floats
	list := data.List(symbol)
	var values []float64

	if len(list) < a.period {
		return false, nil
	}

	for i := 0; i < len(list); i++ {
		values = append(values, list[i].Price())
	}

	// calculate EMA
	a.ema = ta.Mean(values)

	a.values = ema(values, a.period, false)

	// save the calculated ema to the event metrics
	event.Add(fmt.Sprintf("EMA%d", a.period), a.ema)

	return true, nil
}

// Value returns the value of this Algo.
func (a *emaAlgo) Value() interface{} {
	return a.values
}

func ema(in []float64, period int, macd bool) []float64 {
	var out []float64
	if !macd {
		out = make([]float64, len(in))
	}
	if len(in) < period {
		return out
	}

	smaRet := sma(in, period, macd)
	if macd {
		out = append(out, smaRet[0])
	} else {
		out[period-1] = smaRet[period-1]
	}
	var multiplier = (2.0 / (float64(period) + 1.0))
	for i := period; i < len(in); i++ {
		var lastVal float64
		if macd {
			lastVal = out[len(out)-1]
		} else {
			lastVal = out[i-1]
		}
		ema := (in[i]-lastVal)*multiplier + lastVal
		if macd {
			out = append(out, ema)
			continue
		}
		out[i] = ema
	}
	return out
}

func sma(in []float64, period int, macd bool) []float64 {
	var out []float64
	if !macd {
		out = make([]float64, len(in))
	}
	if len(in) < period {
		return out
	}
	for i := range in {
		if i+1 >= period {
			avg := mean(in[i+1-period : i+1])
			if macd {
				out = append(out, avg)
				continue
			}
			out[i] = avg
		}
	}
	return out
}

func mean(values []float64) float64 {
	var total float64 = 0
	for x := range values {
		total += values[x]
	}
	return total / float64(len(values))
}
