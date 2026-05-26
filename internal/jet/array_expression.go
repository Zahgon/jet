package jet

// Array interface
type Array[E Expression] interface {
	Expression

	EQ(rhs Array[E]) BoolExpression
	NOT_EQ(rhs Array[E]) BoolExpression
	LT(rhs Array[E]) BoolExpression
	GT(rhs Array[E]) BoolExpression
	LT_EQ(rhs Array[E]) BoolExpression
	GT_EQ(rhs Array[E]) BoolExpression

	CONTAINS(rhs Array[E]) BoolExpression
	IS_CONTAINED_BY(rhs Array[E]) BoolExpression
	OVERLAP(rhs Array[E]) BoolExpression
	CONCAT(rhs Array[E]) Array[E]
	CONCAT_ELEMENT(E) Array[E]

	AT(expression IntegerExpression) E
}

type arrayInterfaceImpl[E Expression] struct {
	parent Array[E]
}

type BinaryBoolOp func(Expression, Expression) BoolExpression

func (a arrayInterfaceImpl[E]) EQ(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) NOT_EQ(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) LT(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) GT(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) LT_EQ(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) GT_EQ(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) CONTAINS(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) IS_CONTAINED_BY(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) OVERLAP(rhs Array[E]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (a arrayInterfaceImpl[E]) CONCAT(rhs Array[E]) Array[E] { _ = "STUB: not implemented"; return nil }

func (a arrayInterfaceImpl[E]) CONCAT_ELEMENT(rhs E) Array[E] {
	_ = "STUB: not implemented"
	return nil
}

func (a arrayInterfaceImpl[E]) AT(at IntegerExpression) E {
	_ = "STUB: not implemented"
	return *new(E)
}

type arrayExpressionWrapper[E Expression] struct {
	arrayInterfaceImpl[E]
	Expression
}

func newArrayExpressionWrap[E Expression](expression Expression) Array[E] {
	_ = "STUB: not implemented"
	return nil
}

// ArrayExp is array expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as array expression.
// Does not add sql cast to generated sql builder output.
func ArrayExp[E Expression](expression Expression) Array[E] { _ = "STUB: not implemented"; return nil }

// CastToArrayElemType casts exp to array element type
func CastToArrayElemType[E Expression](array Array[E], exp Expression) E {
	_ = "STUB: not implemented"
	return *new(E)
}

// ARRAY constructor builds an array value using list of expressions.
func ARRAY[E Expression](elems ...E) Array[E] { _ = "STUB: not implemented"; return nil }
