package jet

// TimeExpression interface
type TimeExpression interface {
	Expression

	EQ(rhs TimeExpression) BoolExpression
	NOT_EQ(rhs TimeExpression) BoolExpression
	IS_DISTINCT_FROM(rhs TimeExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs TimeExpression) BoolExpression

	LT(rhs TimeExpression) BoolExpression
	LT_EQ(rhs TimeExpression) BoolExpression
	GT(rhs TimeExpression) BoolExpression
	GT_EQ(rhs TimeExpression) BoolExpression
	BETWEEN(min, max TimeExpression) BoolExpression
	NOT_BETWEEN(min, max TimeExpression) BoolExpression

	ADD(rhs Interval) TimeExpression
	SUB(rhs Interval) TimeExpression
}

type timeInterfaceImpl struct {
	root TimeExpression
}

func (t *timeInterfaceImpl) EQ(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) NOT_EQ(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) IS_DISTINCT_FROM(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) LT(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) LT_EQ(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) GT(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) GT_EQ(rhs TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) BETWEEN(min, max TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) NOT_BETWEEN(min, max TimeExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timeInterfaceImpl) ADD(rhs Interval) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

func (t *timeInterfaceImpl) SUB(rhs Interval) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

//---------------------------------------------------//

type timeExpressionWrapper struct {
	Expression
	timeInterfaceImpl
}

func newTimeExpressionWrap(expression Expression) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

// TimeExp is time expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as time expression.
// Does not add sql cast to generated sql builder output.
func TimeExp(expression Expression) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}
