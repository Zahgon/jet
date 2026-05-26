package jet

// RowExpression interface
type RowExpression interface {
	Expression
	HasProjections

	EQ(rhs RowExpression) BoolExpression
	NOT_EQ(rhs RowExpression) BoolExpression
	IS_DISTINCT_FROM(rhs RowExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs RowExpression) BoolExpression

	LT(rhs RowExpression) BoolExpression
	LT_EQ(rhs RowExpression) BoolExpression
	GT(rhs RowExpression) BoolExpression
	GT_EQ(rhs RowExpression) BoolExpression
}

type rowInterfaceImpl struct {
	root        Expression
	dialect     Dialect
	expressions []Expression
}

func (n *rowInterfaceImpl) EQ(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) NOT_EQ(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) IS_DISTINCT_FROM(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) GT(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) GT_EQ(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) LT(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) LT_EQ(rhs RowExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

func (n *rowInterfaceImpl) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

// ---------------------------------------------------//
type rowExpressionWrapper struct {
	rowInterfaceImpl
	Expression
}

func newRowExpression(name string, dialect Dialect, expressions ...Expression) RowExpression {
	_ = "STUB: not implemented"
	return *new(RowExpression)
}

// ROW function is used to create a tuple value that consists of a set of expressions or column values.
func ROW(dialect Dialect, expressions ...Expression) RowExpression {
	_ = "STUB: not implemented"
	return *new(RowExpression)
}

// WRAP creates row expressions without ROW keyword `( expression1, expression2, ... )`.
func WRAP(dialect Dialect, expressions ...Expression) RowExpression {
	_ = "STUB: not implemented"
	return *new(RowExpression)
}

// RowExp serves as a wrapper for an arbitrary expression, treating it as a row expression.
// This enables the Go compiler to interpret any expression as a row expression
// Note: This does not modify the generated SQL builder output by adding a SQL CAST operation.
func RowExp(expression Expression) RowExpression {
	_ = "STUB: not implemented"
	return *new(RowExpression)
}
