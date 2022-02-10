package indicator

import gbt "github.com/dirkolbrich/gobacktest"

type fixedAlgo struct {
	gbt.Algo
	value float64
}

func Fixed(fixedValue float64) gbt.AlgoHandler {
	return &fixedAlgo{value: fixedValue}
}

func (f *fixedAlgo) Run(s gbt.StrategyHandler) (bool, error) {
	return true, nil
}

func (f *fixedAlgo) Value() interface{} {
	return f.value
}
