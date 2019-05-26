package totp_test

import (
	"crypto/sha1"
	"testing"
	"totp"
	"totp/util"
)

func TestHOTP_ComputeString(t *testing.T) {
	// https://tools.ietf.org/html/rfc4226#page-32
	testdata := []*struct {
		Counter int
		Expected string
	}{
		{0, "755224"},
		{1, "287082"},
		{2, "359152"},
		{3, "969429"},
		{4, "338314"},
		{5, "254676"},
		{6, "287922"},
		{7, "162583"},
		{8, "399871"},
		{9, "520489"},
	}

	hotpImpl := totp.NewHOTP(sha1.New, []byte("12345678901234567890"))

	for _, testcase := range testdata {
		actual := hotpImpl.ComputeString(util.IntBytes(uint64(testcase.Counter)), 6)

		if actual != testcase.Expected {
			t.Error("HOTP: invalid result", actual, testcase.Expected)
		}
	}
}
