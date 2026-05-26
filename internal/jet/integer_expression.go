package jet

// IntegerExpression interface
type IntegerExpression interface {
	Expression
	numericExpression

	EQ(rhs IntegerExpression) BoolExpression
	NOT_EQ(rhs IntegerExpression) BoolExpression
	IS_DISTINCT_FROM(rhs IntegerExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs IntegerExpression) BoolExpression

	LT(rhs IntegerExpression) BoolExpression
	LT_EQ(rhs IntegerExpression) BoolExpression
	GT(rhs IntegerExpression) BoolExpression
	GT_EQ(rhs IntegerExpression) BoolExpression
	BETWEEN(min, max IntegerExpression) BoolExpression
	NOT_BETWEEN(min, max IntegerExpression) BoolExpression

	ADD(rhs IntegerExpression) IntegerExpression
	SUB(rhs IntegerExpression) IntegerExpression
	MUL(rhs IntegerExpression) IntegerExpression
	DIV(rhs IntegerExpression) IntegerExpression
	MOD(rhs IntegerExpression) IntegerExpression
	POW(rhs IntegerExpression) IntegerExpression

	BIT_AND(rhs IntegerExpression) IntegerExpression
	BIT_OR(rhs IntegerExpression) IntegerExpression
	BIT_XOR(rhs IntegerExpression) IntegerExpression
	BIT_SHIFT_LEFT(shift IntegerExpression) IntegerExpression
	BIT_SHIFT_RIGHT(shift IntegerExpression) IntegerExpression
}

// additional integer expression subtypes, used in range expressions.
type (
	Int4Expression IntegerExpression
	Int8Expression IntegerExpression
)

type integerInterfaceImpl struct {
	numericExpressionImpl
	root IntegerExpression
}

func (i *integerInterfaceImpl) EQ(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) NOT_EQ(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) IS_DISTINCT_FROM(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) GT(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) GT_EQ(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) LT(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) LT_EQ(rhs IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) BETWEEN(min, max IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) NOT_BETWEEN(min, max IntegerExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (i *integerInterfaceImpl) ADD(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) SUB(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) MUL(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) DIV(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) MOD(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) POW(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) BIT_AND(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) BIT_OR(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) BIT_XOR(rhs IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) BIT_SHIFT_LEFT(intExpression IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func (i *integerInterfaceImpl) BIT_SHIFT_RIGHT(intExpression IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func newBinaryIntegerOperatorExpression(lhs, rhs IntegerExpression, operator string) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

func newPrefixIntegerOperatorExpression(expression IntegerExpression, operator string) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

type integerExpressionWrapper struct {
	integerInterfaceImpl

	Expression
}

func newIntExpressionWrap(expression Expression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// IntExp is int expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as int expression.
// Does not add sql cast to generated sql builder output.
func IntExp(expression Expression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}
