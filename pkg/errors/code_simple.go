package errors

import "fmt"

// Code fields: [Class][ErrType]
// Tree-like structure
// Class: api
//    ErrType: unknown
//    ErrType: validation error
//    ErrType: authorization error
// Class: jobs
//    ErrType: unknown
//    ErrType: database query
//    ErrType: timeout

type ClassCode uint8
type SimpleErrorCode uint8

type SimpleCode struct {
	Class   ClassCode
	ErrType SimpleErrorCode
}

type SimpleErrorInfo struct {
	Value       SimpleErrorCode
	Name        string
	Description string
}

type SimpleClassInfo struct {
	Value       ClassCode
	Name        string
	Description string
	ErrorTypes  []SimpleErrorInfo
}

var SimpleCodeTree = []SimpleClassInfo{
	{
		Value:       0,
		Name:        "unknown",
		Description: "Unknown",
		ErrorTypes: []SimpleErrorInfo{
			{
				Value:       0,
				Name:        "unknown",
				Description: "Unknown API error",
			},
		},
	},
	{
		Value:       1,
		Name:        "api",
		Description: "API related errors",
		ErrorTypes: []SimpleErrorInfo{
			{
				Value:       0,
				Name:        "unknown",
				Description: "Unknown API error",
			},
			{
				Value:       1,
				Name:        "validation_error",
				Description: "API validation error",
			},
			{
				Value:       2,
				Name:        "authorization_error",
				Description: "API authorization error",
			},
		},
	},
	{
		Value:       2,
		Name:        "jobs",
		Description: "Background job errors",
		ErrorTypes: []SimpleErrorInfo{
			{
				Value:       0,
				Name:        "unknown",
				Description: "Unknown job error",
			},
			{
				Value:       1,
				Name:        "database_query",
				Description: "Database query error in job",
			},
			{
				Value:       2,
				Name:        "timeout",
				Description: "Job execution timeout",
			},
		},
	},
	{
		Value:       255,
		Name:        "max",
		Description: "Example max value",
		ErrorTypes: []SimpleErrorInfo{
			{
				Value:       255,
				Name:        "max",
				Description: "Max error type number",
			},
		},
	},
}

func (SimpleCode) GetType() CodeType {
	return CodeTypeSimple
}

func (e SimpleCode) Encode() string {
	// Pack into 16 bits:
	// - 8 bits for Class
	// - 8 bits for ErrorType
	packed := uint32(e.Class)<<8 | uint32(e.ErrType)

	// Convert type and data to base36
	typeStr := toBase36(uint32(e.GetType()), 1) // 1 char for type
	dataStr := toBase36(packed, 4)              // 4 chars for 16 bits of data (increased from 3)

	return fmt.Sprintf("E%s%s", typeStr, dataStr)
}

func DecodeSimpleCode(code string) (SimpleCode, error) {
	if len(code) != 6 || code[0] != 'E' { // Format: E<type><data> (length now 6 due to 4 chars data)
		return SimpleCode{}, fmt.Errorf("invalid code format: %s", code)
	}

	// Parse type
	typeVal := CodeType(fromBase36(code[1:2]))
	if typeVal != CodeTypeSimple {
		return SimpleCode{}, fmt.Errorf("invalid code type: %d", typeVal)
	}

	// Parse data
	packed := fromBase36(code[2:])

	return SimpleCode{
		Class:   ClassCode((packed >> 8) & 0xFF),
		ErrType: SimpleErrorCode(packed & 0xFF),
	}, nil
}

// String returns a human-readable representation of the error code
func (e SimpleCode) String() string {
	for _, class := range SimpleCodeTree {
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

func (SimpleCode) GetPrefix() string {
	return "E"
}

func (SimpleCode) GetFieldInfo() []FieldInfo {
	return []FieldInfo{
		{
			Name:        "Class",
			Bits:        8,
			Description: "Identifies the error class (0-255)",
			Values:      "api(1), jobs(2)",
		},
		{
			Name:        "ErrorType",
			Bits:        8,
			Description: "Identifies the specific error (0-255)",
			Values:      "unknown(0), validation(1), etc",
		},
	}
}

func (SimpleCode) GetDocSection() DocSection {
	return DocSection{
		Title: "Simple Format",
		Description: `Each error code is composed of two bytes encoded as follows:
- Class (8 bits): Identifies the error class (allows up to 256 distinct classes)
- ErrorType (8 bits): Identifies the specific error (allows up to 256 errors per class)

The format provides:
- Up to 256 different classes
- Up to 256 different error types per class
- Total of 65,536 possible unique error codes

The code is encoded as E<type><data> where:
- E: Fixed prefix
- type: 1 character in base-36 encoding the error type
- data: 4 characters in base-36 encoding the class and error type bits

Bit layout before encoding:
` + "```" + `
[CCCCCCCC][EEEEEEEE]
C: Class bits
E: ErrorType bits
` + "```",
		Headers: []string{"Code", "Class.Type", "Description"},
	}
}

// GetPermutations returns all possible error code combinations
func (SimpleCode) GetPermutations() []Permutation {
	var perms []Permutation

	for _, class := range SimpleCodeTree {
		for _, errType := range class.ErrorTypes {
			sc := SimpleCode{
				Class:   class.Value,
				ErrType: errType.Value,
			}
			code := sc.Encode()

			// Validate by decoding and comparing
			decoded, err := DecodeSimpleCode(code)
			if err != nil {
				panic(fmt.Sprintf("failed to decode simple code: %v", err))
			}
			if decoded.Class != sc.Class || decoded.ErrType != sc.ErrType {
				panic(fmt.Sprintf("decode mismatch: got %v, want %v", decoded, sc))
			}

			perms = append(perms, Permutation{
				Type: CodeTypeSimple,
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
