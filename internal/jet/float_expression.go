package jet

// FloatExpression is interface for SQL float columns
type FloatExpression interface {
	Expression
	numericExpression

	EQ(rhs FloatExpression) BoolExpression
	NOT_EQ(rhs FloatExpression) BoolExpression
	IS_DISTINCT_FROM(rhs FloatExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs FloatExpression) BoolExpression

	LT(rhs FloatExpression) BoolExpression
	LT_EQ(rhs FloatExpression) BoolExpression
	GT(rhs FloatExpression) BoolExpression
	GT_EQ(rhs FloatExpression) BoolExpression
	BETWEEN(min, max FloatExpression) BoolExpression
	NOT_BETWEEN(min, max FloatExpression) BoolExpression

	ADD(rhs NumericExpression) FloatExpression
	SUB(rhs NumericExpression) FloatExpression
	MUL(rhs NumericExpression) FloatExpression
	DIV(rhs NumericExpression) FloatExpression
	MOD(rhs NumericExpression) FloatExpression
	POW(rhs NumericExpression) FloatExpression
}

type floatInterfaceImpl struct {
	numericExpressionImpl
	root FloatExpression
}

func (n *floatInterfaceImpl) EQ(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) NOT_EQ(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) IS_DISTINCT_FROM(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) GT(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) GT_EQ(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) LT(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) LT_EQ(rhs FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) BETWEEN(min, max FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) NOT_BETWEEN(min, max FloatExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *floatInterfaceImpl) ADD(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (n *floatInterfaceImpl) SUB(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (n *floatInterfaceImpl) MUL(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (n *floatInterfaceImpl) DIV(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (n *floatInterfaceImpl) MOD(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

func (n *floatInterfaceImpl) POW(rhs NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *

	//---------------------------------------------------//
	new(FloatExpression)
}

type floatExpressionWrapper struct {
	floatInterfaceImpl
	Expression
}

func newFloatExpressionWrap(expression Expression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// FloatExp is date expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as float expression.
// Does not add sql cast to generated sql builder output.
func FloatExp(expression Expression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}
