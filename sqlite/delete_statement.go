package sqlite

import "github.com/go-jet/jet/v2/internal/jet"

// DeleteStatement is interface for MySQL DELETE statement
type DeleteStatement interface {
	Statement

	WHERE(expression BoolExpression) DeleteStatement
	ORDER_BY(orderByClauses ...OrderByClause) DeleteStatement
	LIMIT(limit int64) DeleteStatement
	RETURNING(projections ...Projection) DeleteStatement
}

type deleteStatementImpl struct {
	jet.SerializerStatement

	Delete    jet.ClauseDelete
	Where     jet.ClauseWhere
	OrderBy   jet.ClauseOrderBy
	Limit     jet.ClauseLimit
	Returning jet.ClauseReturning
}

func newDeleteStatement(table Table) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) WHERE(expression BoolExpression) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) ORDER_BY(orderByClauses ...OrderByClause) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) LIMIT(limit int64) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) RETURNING(projections ...jet.Projection) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}
