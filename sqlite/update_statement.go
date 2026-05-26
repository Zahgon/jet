package sqlite

import "github.com/go-jet/jet/v2/internal/jet"

// UpdateStatement is interface of SQL UPDATE statement
type UpdateStatement interface {
	jet.Statement

	SET(value interface{}, values ...interface{}) UpdateStatement
	MODEL(data interface{}) UpdateStatement

	FROM(tables ...ReadableTable) UpdateStatement
	WHERE(expression BoolExpression) UpdateStatement
	RETURNING(projections ...Projection) UpdateStatement
}

type updateStatementImpl struct {
	jet.SerializerStatement

	Update    jet.ClauseUpdate
	From      jet.ClauseFrom
	Set       jet.SetClause
	SetNew    jet.SetClauseNew
	Where     jet.ClauseWhere
	Returning jet.ClauseReturning
}

func newUpdateStatement(table Table, columns []jet.Column) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) SET(value interface{}, values ...interface{}) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) MODEL(data interface{}) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) FROM(tables ...ReadableTable) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) WHERE(expression BoolExpression) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) RETURNING(projections ...Projection) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}
