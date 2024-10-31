package errors

// CodeType identifies the format/structure of the error code
type CodeType uint16

const (
	CodeTypeSimple       CodeType = 1
	CodeTypeAppComponent CodeType = 2
)

type FieldInfo struct {
	Name        string
	Bits        int
	Description string
	Values      string
}

type DocSection struct {
	Title       string
	Description string
	Headers     []string // Column headers for the table
}

type Permutation struct {
	Type        CodeType
	Code        string
	Fields      map[string]string
	TableFields []string // Fields in order for table display
}

// Interface that all error types must implement
type ErrorType interface {
	Encode() string
	String() string
	GetFieldInfo() []FieldInfo
	GetPermutations() []Permutation
	GetType() CodeType
	GetPrefix() string
	GetDocSection() DocSection
}
