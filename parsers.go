package parsers

import "errors"

// Parser - Type declaration for the Parse function
type Parser func(string) (*IdNo, error)

// Global error messages
var (
	ErrInvalidDate        = errors.New("Invalid date")
	ErrInvalidCheckdigit  = errors.New("Invalid checkdigit")
	ErrInvalidFormat      = errors.New("Invalid id format")
	ErrParserDoesNotExist = errors.New("Parser not implemented")
)

var parsers = make(map[string]Parser, 50)

// Constants
const (
	Male   = "Male"
	Female = "Female"
)

// IdNo - Return type for the Parse function
type IdNo struct {
	Valid  bool              `json:"valid"`
	Source string            `json:"source"`
	Data   map[string]string `json:"data"`
}

// New - Create pointer to new IdNo object
func New(id string) *IdNo {
	n := &IdNo{Source: id}
	n.Data = make(map[string]string)
	return n
}

// Register - Register parser function
func Register(name string, parser Parser) {
	parsers[name] = parser
}

// List - Return map of Parsers
func List() map[string]Parser {
	return parsers
}
