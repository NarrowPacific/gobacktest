package indicator

import (
	"fmt"
	"reflect"
	"testing"

	gbt "github.com/dirkolbrich/gobacktest"
	"github.com/dirkolbrich/gobacktest/algo"
)

func TestHighestIntegration(t *testing.T) {
	// set up mock Data Events
	mockdata := algo.TestHelperMockData([]string{
		"2018-07-01",
		"2018-07-02",
		"2018-07-03",
		"2018-07-04",
		"2018-07-05",
	})

	// set close price from 1 to n on mockdata
	for i, data := range mockdata {
		bar := data.(*gbt.Bar)
		bar.Close = float64(i + 1)
		mockdata[i] = bar
	}

	var testCases = []struct {
		msg       string
		mockdata  []gbt.DataEvent
		lookback  int
		source    Source
		runBefore int
		expOk     bool
		expValue  float64
		expErr    error
	}{
		{msg: "test too much data points",
			mockdata:  mockdata,
			lookback:  7,
			source:    CLOSE,
			runBefore: 5,
			expOk:     false,
			expErr:    fmt.Errorf("invalid value length for indicator highest"),
		},
		{msg: "test normal run",
			mockdata:  mockdata,
			lookback:  5,
			source:    CLOSE,
			runBefore: 5,
			expOk:     true,
			expValue:  5,
			expErr:    nil,
		},
	}

	for _, tc := range testCases {
		// set up data handler
		data := &gbt.Data{}
		data.SetStream(tc.mockdata)
		event, _ := data.Next()

		// set up strategy
		strategy := &gbt.Strategy{}
		strategy.SetData(data)
		strategy.SetEvent(event)

		// run the backtest n times to  pull data from stream and fill data.list
		for i := 0; i < tc.runBefore; i++ {
			data.Next()
		}

		// create Algo
		algo := Highest(CLOSE, tc.lookback)

		ok, err := algo.Run(strategy)
		if (ok != tc.expOk) || !reflect.DeepEqual(err, tc.expErr) || algo.Value() != tc.expValue {
			t.Errorf("%v: Lowest(%v): \nexpected %v %v %#v, \nactual   %v %v %#v",
				tc.msg, tc.lookback, tc.expOk, tc.expValue, tc.expErr, ok, algo.Value(), err)
		}
	}

}
