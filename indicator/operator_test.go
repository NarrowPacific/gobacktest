package indicator

import (
	"reflect"
	"testing"

	gbt "github.com/dirkolbrich/gobacktest"
)

func TestOperatorIntegration(t *testing.T) {
	var testCases = []struct {
		msg      string
		operator gbt.AlgoHandler
		expOk    bool
		expValue float64
		expErr   error
	}{
		{
			msg:      "test Multiply",
			operator: Multiply(Fixed(2), Fixed(5)),
			expValue: 10,
			expOk:    true,
			expErr:   nil,
		},
		{
			msg:      "test Multiply",
			operator: Divide(Fixed(4), Fixed(2)),
			expValue: 2,
			expOk:    true,
			expErr:   nil,
		},
		{
			msg:      "test Add",
			operator: Add(Fixed(2), Fixed(5)),
			expValue: 7,
			expOk:    true,
			expErr:   nil,
		},
		{
			msg:      "test Multiply",
			operator: Subtract(Fixed(4), Fixed(2)),
			expValue: 2,
			expOk:    true,
			expErr:   nil,
		},
	}

	for _, tc := range testCases {
		strategy := &gbt.Strategy{}

		operator := tc.operator
		ok, err := operator.Run(strategy)

		if (ok != tc.expOk) || !reflect.DeepEqual(err, tc.expErr) || operator.Value() != tc.expValue {
			t.Errorf("%v: Last: \nexpected %v %v %#v, \nactual   %v %v %#v",
				tc.msg, tc.expOk, tc.expValue, tc.expErr, ok, operator.Value(), err)
		}
	}

}
