package jet

// TimestampzExpression interface
type TimestampzExpression interface {
	Expression

	EQ(rhs TimestampzExpression) BoolExpression
	NOT_EQ(rhs TimestampzExpression) BoolExpression
	IS_DISTINCT_FROM(rhs TimestampzExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs TimestampzExpression) BoolExpression

	LT(rhs TimestampzExpression) BoolExpression
	LT_EQ(rhs TimestampzExpression) BoolExpression
	GT(rhs TimestampzExpression) BoolExpression
	GT_EQ(rhs TimestampzExpression) BoolExpression
	BETWEEN(min, max TimestampzExpression) BoolExpression
	NOT_BETWEEN(min, max TimestampzExpression) BoolExpression

	ADD(rhs Interval) TimestampzExpression
	SUB(rhs Interval) TimestampzExpression
}

type timestampzInterfaceImpl struct {
	root TimestampzExpression
}

func (t *timestampzInterfaceImpl) EQ(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) NOT_EQ(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) IS_DISTINCT_FROM(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) LT(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) LT_EQ(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) GT(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) GT_EQ(rhs TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) BETWEEN(min, max TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) NOT_BETWEEN(min, max TimestampzExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timestampzInterfaceImpl) ADD(rhs Interval) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

func (t *timestampzInterfaceImpl) SUB(rhs Interval) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

//-------------------------------------------------

type timestampzExpressionWrapper struct {
	timestampzInterfaceImpl
	Expression
}

func newTimestampzExpressionWrap(expression Expression) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// TimestampzExp is timestamp with time zone expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp with time zone expression.
// Does not add sql cast to generated sql builder output.
func TimestampzExp(expression Expression) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}
