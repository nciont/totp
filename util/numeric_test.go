package util_test

import (
	"bytes"
	"github.com/nciont/totp/util"
	"math"
	"testing"
)

func TestGetDigits(t *testing.T) {
	testdata := []*struct {
		Number   uint32
		Digits   int
		Expected int
	}{
		{1284755224, 5, 55224},
		{1094287082, 5, 87082},
		{137359152, 6, 359152},
		{1726969429, 6, 969429},
		{1640338314, 8, 40338314},
		{82162583, 8, 82162583},
	}

	for _, testcase := range testdata {
		if actual := util.GetDigits(testcase.Number, testcase.Digits); actual != testcase.Expected {
			t.Log("GetDigits: invalid result", actual, testcase.Expected)
		}
	}
}

func TestIntBytes(t *testing.T) {
	testdata := []*struct {
		Value    uint64
		Expected []byte
	}{
		{0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{math.MaxUint64, []byte{255, 255, 255, 255, 255, 255, 255, 255}},
	}

	for _, testcase := range testdata {
		if actual := util.IntBytes(testcase.Value); !bytes.Equal(actual, testcase.Expected) {
			t.Error("IntBytes: invalid value", actual, testcase.Expected)
		}
	}
}
