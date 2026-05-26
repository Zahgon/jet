package jet

// StringExpression interface
type StringExpression interface {
	Expression
	isStringOrBlob()

	EQ(rhs StringExpression) BoolExpression
	NOT_EQ(rhs StringExpression) BoolExpression
	IS_DISTINCT_FROM(rhs StringExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs StringExpression) BoolExpression

	LT(rhs StringExpression) BoolExpression
	LT_EQ(rhs StringExpression) BoolExpression
	GT(rhs StringExpression) BoolExpression
	GT_EQ(rhs StringExpression) BoolExpression
	BETWEEN(min, max StringExpression) BoolExpression
	NOT_BETWEEN(min, max StringExpression) BoolExpression

	CONCAT(rhs Expression) StringExpression

	LIKE(pattern StringExpression) BoolExpression
	NOT_LIKE(pattern StringExpression) BoolExpression

	REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression
	NOT_REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression
}

type stringInterfaceImpl struct {
	root StringExpression
}

func (s *stringInterfaceImpl) isStringOrBlob() { _ = "STUB: not implemented"; return }

func (s *stringInterfaceImpl) EQ(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) NOT_EQ(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) IS_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) GT(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) GT_EQ(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) LT(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) LT_EQ(rhs StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) BETWEEN(min, max StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) NOT_BETWEEN(min, max StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) CONCAT(rhs Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

func (s *stringInterfaceImpl) LIKE(pattern StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) NOT_LIKE(pattern StringExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (s *stringInterfaceImpl) NOT_REGEXP_LIKE(pattern StringExpression, caseSensitive ...bool) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

type regexpLikeSerializer struct {
	not           bool
	str           StringExpression
	pattern       StringExpression
	caseSensitive bool
}

func (r *regexpLikeSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ---------------------------------------------------//
func newBinaryStringOperatorExpression(lhs, rhs Expression, operator string) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

//---------------------------------------------------//

type stringExpressionWrapper struct {
	stringInterfaceImpl
	Expression
}

func newStringExpressionWrap(expression Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// StringExp is string expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as string expression.
// Does not add sql cast to generated sql builder output.
func StringExp(expression Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}
