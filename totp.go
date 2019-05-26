package totp

import (
	"crypto"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"hash"
	"strconv"
	"totp/util"
)

func NewTOTP(secret []byte, hasher func() hash.Hash, stepper func(uint64) uint64) *TOTP {
	return &TOTP{ NewHOTP(hasher, secret), stepper }
}

func NewSHA1TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, crypto.SHA1.New, util.TimeStepper)
}

func NewSHA256TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, crypto.SHA256.New, util.TimeStepper)
}

func NewSHA512TOTP(secret []byte) *TOTP {
	return NewTOTP(secret, crypto.SHA512.New, util.TimeStepper)
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
