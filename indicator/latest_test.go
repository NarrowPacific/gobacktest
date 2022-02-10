package indicator

import (
	"reflect"
	"testing"

	gbt "github.com/dirkolbrich/gobacktest"
	"github.com/dirkolbrich/gobacktest/algo"
)

func TestLastIntegration(t *testing.T) {
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
		source    Source
		runBefore int
		expOk     bool
		expValue  float64
		expErr    error
	}{
		{msg: "test normal run",
			mockdata:  mockdata,
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

		// run the backtest n times to pull data from stream and fill data.list
		for i := 0; i < tc.runBefore; i++ {
			event, ok := data.Next()
			if ok {
				strategy.SetEvent(event)
			}
		}

		// create Algo
		algo := Latest(CLOSE)

		ok, err := algo.Run(strategy)
		if (ok != tc.expOk) || !reflect.DeepEqual(err, tc.expErr) || algo.Value() != tc.expValue {
			t.Errorf("%v: Last: \nexpected %v %v %#v, \nactual   %v %v %#v",
				tc.msg, tc.expOk, tc.expValue, tc.expErr, ok, algo.Value(), err)
		}
	}

}
