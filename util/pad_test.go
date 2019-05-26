package util_test

import (
	"github.com/nciont/totp/util"
	"testing"
)

func TestPadStart(t *testing.T) {
	testcases := []*struct{
		Value string
		Length int
		Expected string
	}{
		{ "1234", 8, "00001234" },
		{ "1234", 3, "1234" },
		{ "", 4, "0000" },
	}

	for _, testcase := range testcases {
		if actual := util.PadStart(testcase.Value, "0", testcase.Length); actual != testcase.Expected {
			t.Error("PadStart: invalid result", actual, testcase.Expected)
		}
	}
}
