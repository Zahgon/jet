package jet

// WITH function creates new with statement from list of common table expressions for specified dialect
func WITH(dialect Dialect, recursive bool, cte ...*CommonTableExpression) func(statement Statement) Statement {
	_ = "STUB: not implemented"
	return nil
}

type withImpl struct {
	statementInterfaceImpl
	recursive        bool
	ctes             []*CommonTableExpression
	primaryStatement SerializerStatement
}

func (w withImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func (w withImpl) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *

	// CommonTableExpression contains information about a CTE.
	new(ProjectionList)
}

type CommonTableExpression struct {
	selectTableImpl

	NotMaterialized bool
	Columns         []ColumnExpression
}

// CTE creates new named CommonTableExpression
func CTE(name string, columns ...ColumnExpression) CommonTableExpression {
	_ = "STUB: not implemented"
	return *new(CommonTableExpression)
}

func (c CommonTableExpression) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// serialize CTE definition

// serialize CTE in FROM clause
