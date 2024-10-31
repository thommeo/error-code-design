package errors

import "fmt"

// Code fields: [Class(5)][ErrType(11)]
// First byte:  [CCCCCEEE]
// Second byte: [EEEEEEEE]
// Allows for 32 classes and 2048 error types per class

type Class5Code uint8
type Simple11ErrorCode uint16

type Simple511Code struct {
	Class   Class5Code
	ErrType Simple11ErrorCode
}

type Simple11ErrorInfo struct {
	Value       Simple11ErrorCode
	Name        string
	Description string
}

type Simple5ClassInfo struct {
	Value       Class5Code
	Name        string
	Description string
	ErrorTypes  []Simple11ErrorInfo
}

// Example error codes - you can expand this based on your needs
var Simple511CodeTree = []Simple5ClassInfo{
	{
		Value:       0,
		Name:        "unknown",
		Description: "Unknown error class",
		ErrorTypes: []Simple11ErrorInfo{
			{
				Value:       0,
				Name:        "unknown",
				Description: "Unknown error",
			},
		},
	},
	{
		Value:       1,
		Name:        "http",
		Description: "HTTP-related errors",
		ErrorTypes: []Simple11ErrorInfo{
			{
				Value:       0,
				Name:        "unknown",
				Description: "Unknown HTTP error",
			},
			{
				Value:       1,
				Name:        "bad_request",
				Description: "Bad request error (400)",
			},
			{
				Value:       2,
				Name:        "unauthorized",
				Description: "Unauthorized error (401)",
			},
			{
				Value:       3,
				Name:        "forbidden",
				Description: "Forbidden error (403)",
			},
			{
				Value:       4,
				Name:        "not_found",
				Description: "Not found error (404)",
			},
		},
	},
	{
		Value:       31,
		Name:        "max",
		Description: "Maximum class value example",
		ErrorTypes: []Simple11ErrorInfo{
			{
				Value:       2047,
				Name:        "max",
				Description: "Maximum error type value",
			},
		},
	},
}

func (Simple511Code) GetType() CodeType {
	return CodeType(CodeTypeSimple511)
}

func (e Simple511Code) Encode() string {
	// Pack into 16 bits:
	// - 5 bits for Class
	// - 11 bits for ErrorType
	// Resulting in two bytes: [CCCCCEEE][EEEEEEEE]
	packed := uint32(e.Class)<<11 | uint32(e.ErrType)

	// Convert type and data to base36
	typeStr := toBase36(uint32(e.GetType()), 1) // 1 char for type
	dataStr := toBase36(packed, 4)              // 4 chars for 16 bits of data

	return fmt.Sprintf("E%s%s", typeStr, dataStr)
}

func DecodeSimple511Code(code string) (Simple511Code, error) {
	if len(code) != 6 || code[0] != 'E' { // Format: E<type><data>
		return Simple511Code{}, fmt.Errorf("invalid code format: %s", code)
	}

	// Parse type
	typeVal := CodeType(fromBase36(code[1:2]))
	if typeVal != CodeType(3) {
		return Simple511Code{}, fmt.Errorf("invalid code type: %d", typeVal)
	}

	// Parse data
	packed := fromBase36(code[2:])

	// Validate packed data doesn't exceed our bit limits
	if packed >= 1<<16 {
		return Simple511Code{}, fmt.Errorf("packed value exceeds 16 bits: %d", packed)
	}

	return Simple511Code{
		Class:   Class5Code((packed >> 11) & 0x1F), // 5 bits
		ErrType: Simple11ErrorCode(packed & 0x7FF), // 11 bits
	}, nil
}

func (e Simple511Code) String() string {
	for _, class := range Simple511CodeTree {
		if class.Value != e.Class {
			continue
		}
		for _, errType := range class.ErrorTypes {
			if errType.Value != e.ErrType {
				continue
			}
			return fmt.Sprintf("%s.%s", class.Name, errType.Name)
		}
	}
	return "invalid"
}

func (Simple511Code) GetPrefix() string {
	return "E"
}

func (Simple511Code) GetFieldInfo() []FieldInfo {
	return []FieldInfo{
		{
			Name:        "Class",
			Bits:        5,
			Description: "Identifies the error class (0-31)",
			Values:      "http(1), max(31)",
		},
		{
			Name:        "ErrorType",
			Bits:        11,
			Description: "Identifies the specific error (0-2047)",
			Values:      "unknown(0), bad_request(1), unauthorized(2), etc",
		},
	}
}

func (Simple511Code) GetDocSection() DocSection {
	return DocSection{
		Title: "Simple 5-11 Format",
		Description: `Each error code is composed of 16 bits encoded as follows:
- Class (5 bits): Identifies the error class (allows up to 32 distinct classes)
- ErrorType (11 bits): Identifies the specific error (allows up to 2048 errors per class)

The format provides:
- Up to 32 different classes
- Up to 2048 different error types per class
- Total of 65,536 possible unique error codes

The code is encoded as E<type><data> where:
- E: Fixed prefix
- type: 1 character in base-36 encoding the error type
- data: 4 characters in base-36 encoding the class and error type bits

Bit layout before encoding:
` + "```" + `
[CCCCCEEE][EEEEEEEE]
C: Class bits (5)
E: ErrorType bits (11)
` + "```",
		Headers: []string{"Code", "Class.Type", "Description"},
	}
}

func (Simple511Code) GetPermutations() []Permutation {
	var perms []Permutation

	for _, class := range Simple511CodeTree {
		for _, errType := range class.ErrorTypes {
			sc := Simple511Code{
				Class:   class.Value,
				ErrType: errType.Value,
			}
			code := sc.Encode()

			// Validate by decoding and comparing
			decoded, err := DecodeSimple511Code(code)
			if err != nil {
				panic(fmt.Sprintf("failed to decode simple 5-11 code: %v", err))
			}
			if decoded.Class != sc.Class || decoded.ErrType != sc.ErrType {
				panic(fmt.Sprintf("decode mismatch: got %v, want %v", decoded, sc))
			}

			perms = append(perms, Permutation{
				Type: CodeType(3),
				Code: code,
				Fields: map[string]string{
					"Class":       class.Name,
					"ErrorType":   errType.Name,
					"Description": errType.Description,
				},
				TableFields: []string{
					code,
					fmt.Sprintf("%s.%s", class.Name, errType.Name),
					errType.Description,
				},
			})
		}
	}
	return perms
}
