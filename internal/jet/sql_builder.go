package jet

import (
	"bytes"
)

// SQLBuilder generates output SQL
type SQLBuilder struct {
	Dialect Dialect
	Buff    bytes.Buffer
	Args    []interface{}

	lastChar byte
	ident    int

	Debug bool
}

const tabSize = 4
const defaultIdent = 5

// IncreaseIdent adds ident or defaultIdent number of spaces to each new line
func (s *SQLBuilder) IncreaseIdent(ident ...int) { _ = "STUB: not implemented"; return }

// DecreaseIdent removes ident or defaultIdent number of spaces for each new line
func (s *SQLBuilder) DecreaseIdent(ident ...int) { _ = "STUB: not implemented"; return }

// WriteProjections func
func (s *SQLBuilder) WriteProjections(statement StatementType, projections []Projection) {
	_ = "STUB: not implemented"
	return
}

// WriteRowToJsonProjections serializes slice of projections intended for row_to_json json aggregation
func (s *SQLBuilder) WriteRowToJsonProjections(statement StatementType, projections []Projection) {
	_ = "STUB: not implemented"
	return
}

// NewLine adds new line to output SQL
func (s *SQLBuilder) NewLine() { _ = "STUB: not implemented"; return }

func (s *SQLBuilder) write(data []byte) { _ = "STUB: not implemented"; return }

func isPreSeparator(b byte) bool { _ = "STUB: not implemented"; return false }

func isPostSeparator(b byte) bool { _ = "STUB: not implemented"; return false }

// WriteAlias is used to add alias to output SQL
func (s *SQLBuilder) WriteAlias(str string) { _ = "STUB: not implemented"; return }

// WriteString writes sting to output SQL
func (s *SQLBuilder) WriteString(str string) { _ = "STUB: not implemented"; return }

// WriteJsonObjKey serializes json object key
func (s *SQLBuilder) WriteJsonObjKey(key string) { _ = "STUB: not implemented"; return }

// WriteIdentifier adds identifier to output SQL
func (s *SQLBuilder) WriteIdentifier(name string, alwaysQuote ...bool) {
	_ = "STUB: not implemented"
	return
}

func (s *SQLBuilder) shouldQuote(name string, alwaysQuote ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// WriteByte writes byte to output SQL
func (s *SQLBuilder) WriteByte(b byte) { _ = "STUB: not implemented"; return }

func (s *SQLBuilder) finalize() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

func (s *SQLBuilder) insertConstantArgument(arg interface{}) { _ = "STUB: not implemented"; return }

func (s *SQLBuilder) insertParametrizedArgument(arg interface{}) { _ = "STUB: not implemented"; return }

func (s *SQLBuilder) insertRawQuery(raw string, namedArg map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

// one named argument can occur multiple times inside raw string

// if named argument does not exists in raw string do not add argument to the list of arguments
// It can happen if the same argument occurs multiple times in postgres query.

// if placeholder is not unique identifier ($1, $2, etc..), we will replace just one occurrence of the argument
// all occurrences

// just one occurrence

func (s *SQLBuilder) argToString(value interface{}) string { _ = "STUB: not implemented"; return "" }

// If valuer for some reason returns an error, we return error string representation.
// This is fine because argToString is called only from DebugSQL, and DebugSQL shouldn't be used in production.

func integerTypesToString(value interface{}) string { _ = "STUB: not implemented"; return "" }

func shouldQuoteIdentifier(identifier string) bool { _ = "STUB: not implemented"; return false }

// if it is a number we should quote it

// check if contains non ascii characters

func stringQuote(value string) string { _ = "STUB: not implemented"; return "" }
