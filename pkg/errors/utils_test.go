package errors

import "testing"

func TestBase36(t *testing.T) {
	tests := []struct {
		name    string
		value   uint32
		width   int
		encoded string
	}{
		{"zero", 0, 1, "0"},
		{"single digit", 9, 1, "9"},
		{"letter", 10, 1, "A"},
		{"max letter", 35, 1, "Z"},
		{"two digits", 36, 2, "10"},
		{"mixed", 71, 2, "1Z"},
		{"padding", 5, 3, "005"},

		// Some full 24-bit values
		{"24bit-zero", 0, 5, "00000"},
		{"24bit-max", 0xFFFFFF, 5, "9ZLDR"}, // 16777215
		{"24bit-mid", 0x7FFFFF, 5, "4ZSOV"}, // 8388607
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test encoding
			encoded := toBase36(tt.value, tt.width)
			if encoded != tt.encoded {
				t.Errorf("toBase36(%d, %d) = %s; want %s",
					tt.value, tt.width, encoded, tt.encoded)
			}

			// Test decoding
			decoded := fromBase36(tt.encoded)
			if decoded != tt.value {
				t.Errorf("fromBase36(%s) = %d; want %d",
					tt.encoded, decoded, tt.value)
			}
		})
	}
}

func TestBase36Invalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"lowercase", "abc"},
		{"symbols", "!@#"},
		{"spaces", "1 2"},
		{"unicode", "⌘"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("fromBase36(%s) should panic", tt.input)
				}
			}()
			fromBase36(tt.input)
		})
	}
}
