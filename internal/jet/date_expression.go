package jet

// DateExpression is interface for date types
type DateExpression interface {
	Expression

	EQ(rhs DateExpression) BoolExpression
	NOT_EQ(rhs DateExpression) BoolExpression
	IS_DISTINCT_FROM(rhs DateExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs DateExpression) BoolExpression

	LT(rhs DateExpression) BoolExpression
	LT_EQ(rhs DateExpression) BoolExpression
	GT(rhs DateExpression) BoolExpression
	GT_EQ(rhs DateExpression) BoolExpression
	BETWEEN(min, max DateExpression) BoolExpression
	NOT_BETWEEN(min, max DateExpression) BoolExpression

	ADD(rhs Interval) TimestampExpression
	SUB(rhs Interval) TimestampExpression
}

type dateInterfaceImpl struct {
	root DateExpression
}

func (d *dateInterfaceImpl) EQ(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) NOT_EQ(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) IS_DISTINCT_FROM(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) LT(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) LT_EQ(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) GT(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) GT_EQ(rhs DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) BETWEEN(min, max DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) NOT_BETWEEN(min, max DateExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (d *dateInterfaceImpl) ADD(rhs Interval) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

func (d *dateInterfaceImpl) SUB(rhs Interval) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

//---------------------------------------------------//

type dateExpressionWrapper struct {
	dateInterfaceImpl
	Expression
}

func newDateExpressionWrap(expression Expression) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}

// DateExp is date expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as date expression.
// Does not add sql cast to generated sql builder output.
func DateExp(expression Expression) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}
