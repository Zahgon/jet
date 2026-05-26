package mysql

import "github.com/go-jet/jet/v2/internal/jet"

// UpdateStatement is interface of SQL UPDATE statement
type UpdateStatement interface {
	jet.Statement

	OPTIMIZER_HINTS(hints ...OptimizerHint) UpdateStatement

	SET(value interface{}, values ...interface{}) UpdateStatement
	MODEL(data interface{}) UpdateStatement

	WHERE(expression BoolExpression) UpdateStatement
	LIMIT(limit int64) UpdateStatement
}

type updateStatementImpl struct {
	jet.SerializerStatement

	Update jet.ClauseUpdate
	Set    jet.SetClause
	SetNew jet.SetClauseNew
	Where  jet.ClauseWhere
	Limit  jet.ClauseLimit
}

func newUpdateStatement(table Table, columns []jet.Column) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

// Initialize to -1 to indicate no LIMIT

func (u *updateStatementImpl) OPTIMIZER_HINTS(hints ...OptimizerHint) UpdateStatement {
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

func (u *updateStatementImpl) WHERE(expression BoolExpression) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (u *updateStatementImpl) LIMIT(limit int64) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}
