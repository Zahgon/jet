package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// DeleteStatement is interface for PostgreSQL DELETE statement
type DeleteStatement interface {
	jet.SerializerStatement

	USING(tables ...ReadableTable) DeleteStatement
	WHERE(expression BoolExpression) DeleteStatement
	RETURNING(projections ...jet.Projection) DeleteStatement
}

type deleteStatementImpl struct {
	jet.SerializerStatement

	Delete    jet.ClauseDelete
	Using     jet.ClauseFrom
	Where     jet.ClauseWhere
	Returning jet.ClauseReturning
}

func newDeleteStatement(table WritableTable) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) USING(tables ...ReadableTable) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) WHERE(expression BoolExpression) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (d *deleteStatementImpl) RETURNING(projections ...jet.Projection) DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}
