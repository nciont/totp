package util

import "encoding/binary"

func GetDigits(n uint32, digits int) int {
	return int(n) % pow(10, digits)
}

func IntBytes(n uint64) []byte {
	data := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(data, n)
	return data
}

func pow(n, power int) int {
	result := 1

	for power > 0 {
		if power%2 == 1 {
			result *= n
		}

		power >>= 1
		n *= n
	}

	return result
}
