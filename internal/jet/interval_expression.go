package jet

// IntervalExpression interface
type IntervalExpression interface {
	Expression
	isInterval()

	EQ(rhs IntervalExpression) BoolExpression
	NOT_EQ(rhs IntervalExpression) BoolExpression
	IS_DISTINCT_FROM(rhs IntervalExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs IntervalExpression) BoolExpression

	LT(rhs IntervalExpression) BoolExpression
	LT_EQ(rhs IntervalExpression) BoolExpression
	GT(rhs IntervalExpression) BoolExpression
	GT_EQ(rhs IntervalExpression) BoolExpression
	BETWEEN(min, max IntervalExpression) BoolExpression
	NOT_BETWEEN(min, max IntervalExpression) BoolExpression

	ADD(rhs IntervalExpression) IntervalExpression
	SUB(rhs IntervalExpression) IntervalExpression

	MUL(rhs NumericExpression) IntervalExpression
	DIV(rhs NumericExpression) IntervalExpression
}

type intervalInterfaceImpl struct {
	root IntervalExpression
}

func (i *intervalInterfaceImpl) isInterval() { _ = "STUB: not implemented"; return }

func (i *intervalInterfaceImpl) EQ(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) NOT_EQ(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) IS_DISTINCT_FROM(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) LT(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) LT_EQ(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) GT(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) GT_EQ(rhs IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) BETWEEN(min, max IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) NOT_BETWEEN(min, max IntervalExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *intervalInterfaceImpl) ADD(rhs IntervalExpression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

func (i *intervalInterfaceImpl) SUB(rhs IntervalExpression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

func (i *intervalInterfaceImpl) MUL(rhs NumericExpression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

func (i *intervalInterfaceImpl) DIV(rhs NumericExpression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

type intervalWrapper struct {
	intervalInterfaceImpl
	Expression
}

func newIntervalExpressionWrap(expression Expression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

// IntervalExp is interval expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as interval expression.
// Does not add sql cast to generated sql builder output.
func IntervalExp(expression Expression) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

// Interval interface
type Interval interface {
	Serializer
	isInterval()
}
