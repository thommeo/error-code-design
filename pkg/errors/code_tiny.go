package errors

import "fmt"

// TinyCode uses just a simple error type value from 0-1295 (00-ZZ in base-36)
type TinyCode struct {
	ErrType uint16
}

type TinyErrorInfo struct {
	Value       uint16
	Name        string
	Description string
}

// Predefined error codes
var TinyCodeValues = []TinyErrorInfo{
	{
		Value:       0,
		Name:        "unknown",
		Description: "Unknown error",
	},
	{
		Value:       1,
		Name:        "validation",
		Description: "Validation error",
	},
	{
		Value:       2,
		Name:        "not_found",
		Description: "Resource not found",
	},
	{
		Value:       3,
		Name:        "unauthorized",
		Description: "Unauthorized access",
	},
	{
		Value:       4,
		Name:        "bad_request",
		Description: "Bad request",
	},
	{
		Value:       1295,
		Name:        "max",
		Description: "Maximum error value (ZZ)",
	},
}

func (TinyCode) GetType() CodeType {
	return CodeType(CodeTypeTiny)
}

func (e TinyCode) Encode() string {
	if e.ErrType > 1295 { // 35*36 + 35 = maximum value for two base-36 digits
		panic(fmt.Sprintf("error type value %d exceeds maximum of 1295", e.ErrType))
	}

	// Convert type and data to base36
	typeStr := toBase36(uint32(e.GetType()), 1) // always "0"
	dataStr := toBase36(uint32(e.ErrType), 2)   // 2 chars for error type

	return fmt.Sprintf("E%s%s", typeStr, dataStr)
}

func DecodeTinyCode(code string) (TinyCode, error) {
	if len(code) != 4 || code[0] != 'E' { // Format: E0XX
		return TinyCode{}, fmt.Errorf("invalid code format: %s", code)
	}

	// Parse type
	typeVal := CodeType(fromBase36(code[1:2]))
	if typeVal != CodeType(0) {
		return TinyCode{}, fmt.Errorf("invalid code type: %d", typeVal)
	}

	// Parse data
	errType := fromBase36(code[2:])
	if errType > 1295 {
		return TinyCode{}, fmt.Errorf("error type value %d exceeds maximum of 1295", errType)
	}

	return TinyCode{
		ErrType: uint16(errType),
	}, nil
}

func (e TinyCode) String() string {
	for _, info := range TinyCodeValues {
		if info.Value == e.ErrType {
			return info.Name
		}
	}
	return fmt.Sprintf("error_%d", e.ErrType)
}

func (TinyCode) GetPrefix() string {
	return "E"
}

func (TinyCode) GetFieldInfo() []FieldInfo {
	return []FieldInfo{
		{
			Name:        "ErrorType",
			Bits:        11, // Actually only needs 11 bits since max value is 1295
			Description: "Error type value (0-1295)",
			Values:      "unknown(0), validation(1), not_found(2), etc",
		},
	}
}

func (TinyCode) GetDocSection() DocSection {
	return DocSection{
		Title: "Tiny Format",
		Description: `Simplest possible error code format using just an error type value.

The format provides:
- Values from 0 to 1295 (00 to ZZ in base-36)
- Total of 1,296 possible unique error codes

The code is encoded as E0XX where:
- E: Fixed prefix
- 0: Fixed type identifier
- XX: Two base-36 characters encoding the error type (00-ZZ)

Examples:
- E000: Unknown error
- E001: Validation error
- E0ZZ: Maximum value (1295)`,
		Headers: []string{"Code", "Type", "Description"},
	}
}

func (TinyCode) GetPermutations() []Permutation {
	var perms []Permutation

	for _, errType := range TinyCodeValues {
		tc := TinyCode{
			ErrType: errType.Value,
		}
		code := tc.Encode()

		// Validate by decoding and comparing
		decoded, err := DecodeTinyCode(code)
		if err != nil {
			panic(fmt.Sprintf("failed to decode tiny code: %v", err))
		}
		if decoded.ErrType != tc.ErrType {
			panic(fmt.Sprintf("decode mismatch: got %v, want %v", decoded, tc))
		}

		perms = append(perms, Permutation{
			Type: CodeType(0),
			Code: code,
			Fields: map[string]string{
				"ErrorType":   errType.Name,
				"Description": errType.Description,
			},
			TableFields: []string{
				code,
				errType.Name,
				errType.Description,
			},
		})
	}
	return perms
}
