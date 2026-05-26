package jet

// TimestampExpression interface
type TimestampExpression interface {
	Expression

	EQ(rhs TimestampExpression) BoolExpression
	NOT_EQ(rhs TimestampExpression) BoolExpression
	IS_DISTINCT_FROM(rhs TimestampExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs TimestampExpression) BoolExpression

	LT(rhs TimestampExpression) BoolExpression
	LT_EQ(rhs TimestampExpression) BoolExpression
	GT(rhs TimestampExpression) BoolExpression
	GT_EQ(rhs TimestampExpression) BoolExpression
	BETWEEN(min, max TimestampExpression) BoolExpression
	NOT_BETWEEN(min, max TimestampExpression) BoolExpression

	ADD(rhs Interval) TimestampExpression
	SUB(rhs Interval) TimestampExpression
}

type timestampInterfaceImpl struct {
	root TimestampExpression
}

func (t *timestampInterfaceImpl) EQ(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) NOT_EQ(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) IS_DISTINCT_FROM(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) LT(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) LT_EQ(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) GT(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) GT_EQ(rhs TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) BETWEEN(min, max TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) NOT_BETWEEN(min, max TimestampExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampInterfaceImpl) ADD(rhs Interval) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

func (t *timestampInterfaceImpl) SUB(rhs Interval) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

//-------------------------------------------------

type timestampExpressionWrapper struct {
	timestampInterfaceImpl
	Expression
}

func newTimestampExpressionWrap(expression Expression) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// TimestampExp is timestamp expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp expression.
// Does not add sql cast to generated sql builder output.
func TimestampExp(expression Expression) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}
