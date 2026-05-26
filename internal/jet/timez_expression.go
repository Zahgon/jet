package jet

// TimezExpression interface for 'time with time zone' types
type TimezExpression interface {
	Expression

	EQ(rhs TimezExpression) BoolExpression
	NOT_EQ(rhs TimezExpression) BoolExpression
	IS_DISTINCT_FROM(rhs TimezExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs TimezExpression) BoolExpression

	LT(rhs TimezExpression) BoolExpression
	LT_EQ(rhs TimezExpression) BoolExpression
	GT(rhs TimezExpression) BoolExpression
	GT_EQ(rhs TimezExpression) BoolExpression
	BETWEEN(min, max TimezExpression) BoolExpression
	NOT_BETWEEN(min, max TimezExpression) BoolExpression

	ADD(rhs Interval) TimezExpression
	SUB(rhs Interval) TimezExpression
}

type timezInterfaceImpl struct {
	root TimezExpression
}

func (t *timezInterfaceImpl) EQ(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) NOT_EQ(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) IS_DISTINCT_FROM(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) LT(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) LT_EQ(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) GT(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) GT_EQ(rhs TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) BETWEEN(min, max TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) NOT_BETWEEN(min, max TimezExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (t *timezInterfaceImpl) ADD(rhs Interval) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

func (t *timezInterfaceImpl) SUB(rhs Interval) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

//---------------------------------------------------//

type timezExpressionWrapper struct {
	Expression
	timezInterfaceImpl
}

func newTimezExpressionWrap(expression Expression) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// TimezExp is time with time zone expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as time with time zone expression.
// Does not add sql cast to generated sql builder output.
func TimezExp(expression Expression) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}
