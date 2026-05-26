package jet

// BoolExpression interface
type BoolExpression interface {
	Expression

	// Check if this expression is equal to rhs
	EQ(rhs BoolExpression) BoolExpression
	// Check if this expression is not equal to rhs
	NOT_EQ(rhs BoolExpression) BoolExpression
	// Check if this expression is distinct to rhs
	IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression
	// Check if this expression is not distinct to rhs
	IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression

	// Check if this expression is true
	IS_TRUE() BoolExpression
	// Check if this expression is not true
	IS_NOT_TRUE() BoolExpression
	// Check if this expression is false
	IS_FALSE() BoolExpression
	// Check if this expression is not false
	IS_NOT_FALSE() BoolExpression
	// Check if this expression is unknown
	IS_UNKNOWN() BoolExpression
	// Check if this expression is not unknown
	IS_NOT_UNKNOWN() BoolExpression

	// expression AND operator rhs
	AND(rhs BoolExpression) BoolExpression
	// expression OR operator rhs
	OR(rhs BoolExpression) BoolExpression
}

type boolInterfaceImpl struct {
	root BoolExpression
}

func (b *boolInterfaceImpl) EQ(expression BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) NOT_EQ(expression BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) AND(expression BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) OR(expression BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_TRUE() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_NOT_TRUE() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_FALSE() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_NOT_FALSE() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_UNKNOWN() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (b *boolInterfaceImpl) IS_NOT_UNKNOWN() BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func newBinaryBoolOperatorExpression(lhs, rhs Expression, operator string, additionalParams ...Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func newPrefixBoolOperatorExpression(expression Expression, operator string) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func newPostfixBoolOperatorExpression(expression Expression, operator string) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

type boolExpressionWrapper struct {
	boolInterfaceImpl
	Expression
}

func newBoolExpressionWrap(expression Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// BoolExp is bool expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as bool expression.
// Does not add sql cast to generated sql builder output.
func BoolExp(expression Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}
