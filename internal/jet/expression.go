package jet

// Expression is a common interface for all expressions.
// Can be Bool, Int, Float, String, Date, Time, Timez, Timestamp or Timestampz expressions.
type Expression interface {
	Serializer
	Projection
	GroupByClause
	OrderByClause
	expressionOrColumnList

	serializeForJsonValue(statement StatementType, out *SQLBuilder)
	setRoot(root Expression)

	// IS_NULL tests expression whether it is a NULL value.
	IS_NULL() BoolExpression
	// IS_NOT_NULL tests expression whether it is a non-NULL value.
	IS_NOT_NULL() BoolExpression

	// IN checks if this expressions matches any in expressions list
	IN(expressions ...Expression) BoolExpression
	// NOT_IN checks if this expressions is different of all expressions in expressions list
	NOT_IN(expressions ...Expression) BoolExpression

	// AS the temporary alias name to assign to the expression
	AS(alias string) Projection

	// ASC expression will be used to sort query result in ascending order
	ASC() OrderByClause
	// DESC expression will be used to sort query result in descending order
	DESC() OrderByClause
}

// ExpressionInterfaceImpl implements Expression interface methods
type ExpressionInterfaceImpl struct {
	Root Expression
}

func (e *ExpressionInterfaceImpl) isExpressionOrColumnList() { _ = "STUB: not implemented"; return }

func (e *ExpressionInterfaceImpl) setRoot(root Expression) { _ = "STUB: not implemented"; return }

func (e *ExpressionInterfaceImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

// IS_NULL tests expression whether it is a NULL value.
func (e *ExpressionInterfaceImpl) IS_NULL() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// IS_NOT_NULL tests expression whether it is a non-NULL value.
func (e *ExpressionInterfaceImpl) IS_NOT_NULL() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// IN checks if this expressions matches any in expressions list
func (e *ExpressionInterfaceImpl) IN(expressions ...Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// NOT_IN checks if this expressions is different of all expressions in expressions list
func (e *ExpressionInterfaceImpl) NOT_IN(expressions ...Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// AS the temporary alias name to assign to the expression
func (e *ExpressionInterfaceImpl) AS(alias string) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

// ASC expression will be used to sort a query result in ascending order
func (e *ExpressionInterfaceImpl) ASC() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

// DESC expression will be used to sort a query result in descending order
func (e *ExpressionInterfaceImpl) DESC() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

// NULLS_FIRST specifies sort where null values appear before all non-null values
func (e *ExpressionInterfaceImpl) NULLS_FIRST() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

// NULLS_LAST specifies sort where null values appear after all non-null values
func (e *ExpressionInterfaceImpl) NULLS_LAST() OrderByClause {
	_ = "STUB: not implemented"
	return *new(OrderByClause)
}

func (e *ExpressionInterfaceImpl) serializeForGroupBy(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *ExpressionInterfaceImpl) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *ExpressionInterfaceImpl) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *ExpressionInterfaceImpl) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *ExpressionInterfaceImpl) serializeForJsonValue(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *ExpressionInterfaceImpl) serializeForOrderBy(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

type expression struct {
	ExpressionInterfaceImpl
	Serializer
}

func newExpression(serializer Serializer) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// Representation of binary operations (e.g. comparisons, arithmetic)
type binaryOperatorSerializer struct {
	lhs, rhs        Serializer
	additionalParam Serializer
	operator        string
}

func (c *binaryOperatorSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// NewBinaryOperatorExpression creates new binaryOperatorExpression
func NewBinaryOperatorExpression(lhs, rhs Serializer, operator string, additionalParam ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

type serializersWithOperator struct {
	operator    string
	serializers []Serializer
}

func (s *serializersWithOperator) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func newBoolExpressionListOperator(operator string, expressions []BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func newPrefixOperatorExpression(expression Expression, operator string) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func newPostfixOperatorExpression(expression Expression, operator string) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

type betweenOperatorSerializer struct {
	expression Expression
	notBetween bool
	min        Expression
	max        Expression
}

func (b *betweenOperatorSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// NewBetweenOperatorExpression creates new BETWEEN operator expression
func NewBetweenOperatorExpression(expression, min, max Expression, notBetween bool) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}
