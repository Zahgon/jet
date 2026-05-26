package postgres

import "github.com/go-jet/jet/v2/internal/jet"

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
