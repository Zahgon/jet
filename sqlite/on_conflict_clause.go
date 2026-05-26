package sqlite

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

type onConflict interface {
	WHERE(indexPredicate BoolExpression) conflictTarget
	conflictTarget
}

type conflictTarget interface {
	DO_NOTHING() InsertStatement
	DO_UPDATE(action conflictAction) InsertStatement
}

type onConflictClause struct {
	insertStatement  InsertStatement
	indexExpressions []jet.ColumnExpression
	whereClause      jet.ClauseWhere
	do               jet.Serializer
}

func (o *onConflictClause) WHERE(indexPredicate BoolExpression) conflictTarget {
	_ = "STUB: not implemented"
	return *new(conflictTarget)
}

func (o *onConflictClause) DO_NOTHING() InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (o *onConflictClause) DO_UPDATE(action conflictAction) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (o *onConflictClause) Serialize(statementType jet.StatementType, out *jet.SQLBuilder, options ...jet.SerializeOption) {
	_ = "STUB: not implemented"
	return
}

type conflictAction interface {
	jet.Serializer
	WHERE(condition BoolExpression) conflictAction
}

// SET creates conflict action for ON_CONFLICT clause
func SET(assigments ...ColumnAssigment) conflictAction {
	_ = "STUB: not implemented"
	return *new(conflictAction)
}

type updateConflictActionImpl struct {
	jet.Serializer

	doUpdate jet.KeywordClause
	set      jet.SetClauseNew
	where    jet.ClauseWhere
}

func (u *updateConflictActionImpl) WHERE(condition BoolExpression) conflictAction {
	_ = "STUB: not implemented"
	return *new(conflictAction)
}
