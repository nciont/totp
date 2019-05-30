package util

import "strings"

func PadStart(value, pad string, length int) string {
	return strings.Repeat(pad, maxInt(0, length-len(value))) + value
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
