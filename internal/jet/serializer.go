package jet

// SerializeOption type
type SerializeOption int

// Serialize options
const (
	NoWrap SerializeOption = iota
	SkipNewLine
	Ident

	fallTroughOptions // fall trough options

	ShortName
)

// WithFallTrough extends existing serialize options with additional
func (s SerializeOption) WithFallTrough(options []SerializeOption) []SerializeOption {
	_ = "STUB: not implemented"
	return nil
}

// StatementType is type of the SQL statement
type StatementType string

// Statement types
const (
	SelectStatementType        StatementType = "SELECT"
	SelectJsonObjStatementType StatementType = "SELECT_JSON_OBJ"
	SelectJsonArrStatementType StatementType = "SELECT_JSON_ARR"
	InsertStatementType        StatementType = "INSERT"
	UpdateStatementType        StatementType = "UPDATE"
	DeleteStatementType        StatementType = "DELETE"
	SetStatementType           StatementType = "SET"
	LockStatementType          StatementType = "LOCK"
	UnLockStatementType        StatementType = "UNLOCK"
	WithStatementType          StatementType = "WITH"
)

// Serializer interface
type Serializer interface {
	serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption)
}

// Serialize func
func Serialize(exp Serializer, statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func SerializeForOrderBy(exp Expression, statementType StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func contains(options []SerializeOption, option SerializeOption) bool {
	_ = "STUB: not implemented"
	return false
}

// FallTrough filters fall-trough options from the list
func FallTrough(options []SerializeOption) []SerializeOption { _ = "STUB: not implemented"; return nil }

func without(options []SerializeOption, option SerializeOption) []SerializeOption {
	_ = "STUB: not implemented"
	return nil
}

// ListSerializer serializes list of serializers with separator
type ListSerializer struct {
	Serializers []Serializer
	Separator   string
}

func (s ListSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// NewSerializerClauseImpl is constructor for Seralizer with list of clauses
func NewSerializerClauseImpl(clauses ...Clause) Serializer {
	_ = "STUB: not implemented"
	return *new(Serializer)
}

type serializerImpl struct {
	Clauses []Clause
}

func (s serializerImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Token can be used to construct complex custom expressions
type Token string

func (t Token) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// CustomExpression creates new custom expression. When serialized may require parentheses
// depending on context.
func CustomExpression(parts ...Serializer) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// AtomicCustomExpression creates new custom expression. When serialized does not require parentheses.
func AtomicCustomExpression(parts ...Serializer) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

type customSerializer struct {
	parts  []Serializer
	atomic bool
}

func (c *customSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func optionalWrap(out *SQLBuilder, options []SerializeOption, ser func(out *SQLBuilder, options []SerializeOption)) {
	_ = "STUB: not implemented"
	return
}

func wrap(expressions ...Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }
