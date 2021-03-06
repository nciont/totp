package totp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/nciont/totp/util"
	"hash"
	"strconv"
)

func NewTOTP(secret []byte, hasher func() hash.Hash, stepper func(uint64) uint64) *TOTP {
	return &TOTP{NewHOTP(hasher, secret), stepper}
}

func NewSHA1TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, sha1.New, util.TimeStepper)
}

func NewSHA256TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, sha256.New, util.TimeStepper)
}

func NewSHA512TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, sha512.New, util.TimeStepper)
}

type TOTP struct {
	*HOTP
	Stepper func(uint64) uint64
}

func (totp *TOTP) Compute(time uint64, digits int) int {
	totp.hmacImpl.Reset()
	totp.hmacImpl.Write(util.IntBytes(totp.Stepper(time)))
	return util.GetDigits(getCode(totp.hmacImpl.Sum(nil)), digits)
}

func (totp *TOTP) ComputeString(time uint64, digits int) string {
	return util.PadStart(strconv.Itoa(totp.Compute(time, digits)), "0", digits)
}
