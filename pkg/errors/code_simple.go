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
}

func (SimpleCode) GetType() CodeType {
	return CodeTypeSimple
}

// Encode returns the byte representation of the error code
func (e SimpleCode) Encode() []byte {
	result := make([]byte, 3)

	result[0] = byte(e.GetType()) // type byte
	result[1] = byte(e.Class)     // class byte
	result[2] = byte(e.ErrType)   // error type byte

	return result
}

// Decode creates a SimpleCode from a byte slice
func DecodeSimpleCode(data []byte) (SimpleCode, error) {
	if len(data) != 3 {
		return SimpleCode{}, fmt.Errorf("expected 3 bytes, got %d", len(data))
	}

	codeType := CodeType(data[0])
	if codeType != CodeTypeSimple {
		return SimpleCode{}, fmt.Errorf("invalid code type: %d", codeType)
	}

	return SimpleCode{
		Class:   ClassCode(data[1]),
		ErrType: SimpleErrorCode(data[2]),
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

// GetFieldInfo returns metadata about the code fields
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

// Also need to update GetDocSection to include the byte layout diagram
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

Byte layout:
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
			encoded := sc.Encode()

			// Format code with all bytes needing 2 hex chars
			code := fmt.Sprintf("%s%02X%02X%02X",
				sc.GetPrefix(),
				encoded[0], // type byte
				encoded[1], // class byte
				encoded[2]) // error type byte

			// Validate by decoding and comparing
			decoded, err := DecodeSimpleCode(encoded)
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
