package errors

import (
	"fmt"
	"math"
)

const base36Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// toBase36 converts a number to base36 string with fixed width
func toBase36(num uint32, width int) string {
	maxValue := uint32(math.Pow(36, float64(width)))
	if num >= maxValue {
		panic(fmt.Sprintf("number %d too large for width %d (max %d)", num, width, maxValue-1))
	}

	result := make([]byte, width)
	for i := width - 1; i >= 0; i-- {
		result[i] = base36Chars[num%36]
		num /= 36
	}
	return string(result)
}

// fromBase36 converts a base36 string to number
func fromBase36(s string) uint32 {
	if s == "" {
		panic("empty string")
	}

	var result uint32
	for i := 0; i < len(s); i++ {
		c := s[i]
		var val uint32
		switch {
		case c >= '0' && c <= '9':
			val = uint32(c - '0')
		case c >= 'A' && c <= 'Z':
			val = uint32(c - 'A' + 10)
		default:
			panic(fmt.Sprintf("invalid base36 character: %c", c))
		}
		result = result*36 + val
	}
	return result
}
