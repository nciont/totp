package util_test

import (
	"github.com/nciont/totp/util"
	"testing"
)

func TestTimeStepper(t *testing.T) {
	// https://tools.ietf.org/html/rfc6238#section-4.2
	testdata := []*struct{
		Time uint64
		Expected uint64
	}{
		{ 59, 1 },
		{ 60, 2 },
	}

	for _, testcase := range testdata {
		actual := util.TimeStepper(testcase.Time)

		if actual != testcase.Expected {
			t.Error("TimeStepper: invalid result", actual, testcase.Expected)
		}
	}
}
