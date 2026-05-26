package postgres

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

type onConflict interface {
	ON_CONSTRAINT(name string) conflictTarget
	WHERE(indexPredicate BoolExpression) conflictTarget
	conflictTarget
}

type conflictTarget interface {
	DO_NOTHING() InsertStatement
	DO_UPDATE(action conflictAction) InsertStatement
}

type onConflictClause struct {
	insertStatement  InsertStatement
	constraint       string
	indexExpressions []jet.ColumnExpression
	whereClause      jet.ClauseWhere
	do               jet.Serializer
}

func (o *onConflictClause) ON_CONSTRAINT(name string) conflictTarget {
	_ = "STUB: not implemented"
	return *new(conflictTarget)
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
