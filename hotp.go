package totp

import (
	"crypto/hmac"
	"encoding/binary"
	"github.com/nciont/totp/util"
	"hash"
	"strconv"
)

func NewHOTP(hashFn func() hash.Hash, secret []byte) *HOTP {
	return &HOTP{
		hmacImpl: hmac.New(hashFn, secret),
	}
}

type HOTP struct {
	hmacImpl hash.Hash
}

func (hotp *HOTP) Compute(counter []byte, digits int) int {
	hotp.hmacImpl.Reset()
	hotp.hmacImpl.Write(counter)
	return util.GetDigits(getCode(hotp.hmacImpl.Sum(nil)), digits)
}

func (hotp *HOTP) ComputeString(counter []byte, digits int) string {
	return util.PadStart(strconv.Itoa(hotp.Compute(counter, digits)), "0", digits)
}

func getCode(h []byte) uint32 {
	// Optimized version of:
	// offset := h[len(h) - 1] & 0x0f
	// return ((uint32(h[offset]) & 0x7f) << 24)  |
	//     ((uint32(h[offset + 1]) & 0xff) << 16) |
	//	   ((uint32(h[offset + 2]) & 0xff) << 8)  |
	//	   (uint32(h[offset + 3]) & 0xff)
	return binary.BigEndian.Uint32(append(h[:0:0], h[h[len(h) - 1] & 0x0f:]...)) & 0x7fffffff
}
