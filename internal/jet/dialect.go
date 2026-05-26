package jet

// Dialect interface
type Dialect interface {
	Name() string
	PackageName() string
	OperatorSerializeOverride(operator string) SerializeOverride
	AliasQuoteChar() byte
	IdentifierQuoteChar() byte
	ArgumentPlaceholder() QueryPlaceholderFunc
	ArgumentToString(value any) (string, bool)
	IsReservedWord(name string) bool
	SerializeOrderBy() func(expression Expression, ascending, nullsFirst *bool) SerializerFunc
	ValuesDefaultColumnName(index int) string
	JsonValueEncode(expr Expression) Expression
	RegexpLike(str StringExpression, not bool, pattern StringExpression, caseSensitive bool) SerializerFunc
}

// SerializerFunc func
type SerializerFunc func(statement StatementType, out *SQLBuilder, options ...SerializeOption)

// SerializeOverride func
type SerializeOverride func(expressions ...Serializer) SerializerFunc

// QueryPlaceholderFunc func
type QueryPlaceholderFunc func(ord int) string

// DialectParams struct
type DialectParams struct {
	Name                       string
	PackageName                string
	OperatorSerializeOverrides map[string]SerializeOverride
	AliasQuoteChar             byte
	IdentifierQuoteChar        byte
	ArgumentPlaceholder        QueryPlaceholderFunc
	ArgumentToString           func(value any) (string, bool)
	ReservedWords              []string
	SerializeOrderBy           func(expression Expression, ascending, nullsFirst *bool) SerializerFunc
	ValuesDefaultColumnName    func(index int) string
	JsonValueEncode            func(expr Expression) Expression
	RegexpLike                 func(str StringExpression, not bool, pattern StringExpression, caseSensitive bool) SerializerFunc
}

// NewDialect creates new dialect with params
func NewDialect(params DialectParams) Dialect { _ = "STUB: not implemented"; return *new(Dialect) }

type dialectImpl struct {
	name                       string
	packageName                string
	operatorSerializeOverrides map[string]SerializeOverride
	aliasQuoteChar             byte
	identifierQuoteChar        byte
	argumentPlaceholder        QueryPlaceholderFunc
	argumentToString           func(value any) (string, bool)
	reservedWords              map[string]bool
	serializeOrderBy           func(expression Expression, ascending, nullsFirst *bool) SerializerFunc
	valuesDefaultColumnName    func(index int) string
	jsonValueEncode            func(expr Expression) Expression
	regexpLike                 func(str StringExpression, not bool, pattern StringExpression, caseSensitive bool) SerializerFunc
}

func (d *dialectImpl) Name() string { _ = "STUB: not implemented"; return "" }

func (d *dialectImpl) PackageName() string { _ = "STUB: not implemented"; return "" }

func (d *dialectImpl) OperatorSerializeOverride(operator string) SerializeOverride {
	_ = "STUB: not implemented"
	return *new(SerializeOverride)
}

func (d *dialectImpl) AliasQuoteChar() byte { _ = "STUB: not implemented"; return 0 }

func (d *dialectImpl) IdentifierQuoteChar() byte { _ = "STUB: not implemented"; return 0 }

func (d *dialectImpl) ArgumentPlaceholder() QueryPlaceholderFunc {
	_ = "STUB: not implemented"
	return *new(QueryPlaceholderFunc)
}

func (d *dialectImpl) ArgumentToString(value any) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (d *dialectImpl) IsReservedWord(name string) bool { _ = "STUB: not implemented"; return false }

func (d *dialectImpl) SerializeOrderBy() func(expression Expression, ascending, nullsFirst *bool) SerializerFunc {
	_ = "STUB: not implemented"
	return nil
}

func (d *dialectImpl) ValuesDefaultColumnName(index int) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *dialectImpl) JsonValueEncode(expr Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func (d *dialectImpl) RegexpLike(str StringExpression, not bool, pattern StringExpression, caseSensitive bool) SerializerFunc {
	_ = "STUB: not implemented"
	return *new(SerializerFunc)
}

func arrayOfStringsToMapOfStrings(arr []string) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}
