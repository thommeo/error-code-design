package errors

import (
	"fmt"
	"math"
)

const base36Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// validateErrorBits checks if the bits are within their allowed ranges
func validateErrorBits(app, component, subComponent, errType uint32) error {
	if app > 0x0F {
		return fmt.Errorf("app value %d exceeds 4 bits", app)
	}
	if component > 0x3F {
		return fmt.Errorf("component value %d exceeds 6 bits", component)
	}
	if subComponent > 0x3F {
		return fmt.Errorf("subComponent value %d exceeds 6 bits", subComponent)
	}
	if errType > 0xFF {
		return fmt.Errorf("errType value %d exceeds 8 bits", errType)
	}
	return nil
}

// packErrorBits combines the individual fields into a 24-bit value
func packErrorBits(app, component, subComponent, errType uint32) (uint32, error) {
	if err := validateErrorBits(app, component, subComponent, errType); err != nil {
		return 0, err
	}

	return (app << 20) | (component << 14) | (subComponent << 8) | errType, nil
}

// unpackErrorBits extracts individual fields from a 24-bit value
func unpackErrorBits(packed uint32) (app, component, subComponent, errType uint32) {
	app = (packed >> 20) & 0x0F
	component = (packed >> 14) & 0x3F
	subComponent = (packed >> 8) & 0x3F
	errType = packed & 0xFF
	return
}

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

// encodeError encodes error code bits into a base-36 string
func encodeError(app, component, subComponent, errType uint32) (string, error) {
	packed, err := packErrorBits(app, component, subComponent, errType)
	if err != nil {
		return "", err
	}
	return toBase36(packed, 5), nil
}

// decodeError decodes a base-36 string into error code bits
func decodeError(s string) (app, component, subComponent, errType uint32, err error) {
	packed := fromBase36(s)
	if packed >= 1<<24 {
		return 0, 0, 0, 0, fmt.Errorf("value exceeds 24 bits: %d", packed)
	}
	app, component, subComponent, errType = unpackErrorBits(packed)
	return app, component, subComponent, errType, nil
}
