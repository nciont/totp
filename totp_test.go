package totp

import (
	"testing"
)

func TestTOTP(t *testing.T) {
	// https://tools.ietf.org/html/rfc6238#appendix-B
	testdata := []*struct{
		Time uint64
		ExpectedSHA1 string
		ExpectedSHA256 string
		ExpectedSHA512 string
	}{
		{ 59,          "94287082", "46119246", "90693936" },
		{ 1111111109,  "07081804", "68084774", "25091201" },
		{ 1111111111,  "14050471", "67062674", "99943326" },
		{ 1234567890,  "89005924", "91819424", "93441116" },
		{ 2000000000,  "69279037", "90698825", "38618901" },
		{ 20000000000, "65353130", "77737706", "47863826" },
	}

	totpSHA1 := NewSHA1TOTP([]byte("12345678901234567890"))
	totpSHA256 := NewSHA256TOTP([]byte("12345678901234567890123456789012"))
	totpSHA512 := NewSHA512TOTP([]byte("1234567890123456789012345678901234567890123456789012345678901234"))

	var actual string
	for _, testcase := range testdata {
		if actual = totpSHA1.ComputeString(testcase.Time, 8); actual != testcase.ExpectedSHA1 {
			t.Error("totpSHA1: invalid result", actual, testcase.ExpectedSHA1)
		}
		if actual = totpSHA256.ComputeString(testcase.Time, 8); actual != testcase.ExpectedSHA256 {
			t.Error("totpSHA256: invalid result", actual, testcase.ExpectedSHA256)
		}
		if actual = totpSHA512.ComputeString(testcase.Time, 8); actual != testcase.ExpectedSHA512 {
			t.Error("totpSHA512: invalid result", actual, testcase.ExpectedSHA512)
		}
	}
}
