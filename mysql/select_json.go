package mysql

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

// SelectJsonStatement is an interface for MySQL statements that generate JSON on the server.
type SelectJsonStatement interface {
	Statement
	jet.Serializer

	AS(alias string) Projection

	FROM(table ReadableTable) SelectJsonStatement
	WHERE(condition BoolExpression) SelectJsonStatement
	ORDER_BY(orderByClauses ...OrderByClause) SelectJsonStatement
	LIMIT(limit int64) SelectJsonStatement
	OFFSET(offset int64) SelectJsonStatement
}

// SELECT_JSON_ARR creates a new SelectJsonStatement with a list of projections.
func SELECT_JSON_ARR(projections ...Projection) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

// SELECT_JSON_OBJ creates a new SelectJsonStatement with a list of projections.
func SELECT_JSON_OBJ(projections ...Projection) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

type selectJsonStatement struct {
	*selectStatementImpl

	projections   []Projection
	statementType jet.StatementType

	// SELECT_JSON_ARR internal clauses
	arrOrderBy *jet.ClauseOrderBy
	arrLimit   *jet.ClauseLimit
	arrOffset  *jet.ClauseOffset
}

func newSelectStatementJson(projections []Projection, statementType jet.StatementType) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

func (s *selectJsonStatement) constructProjectionList() { _ = "STUB: not implemented"; return }

func (s *selectJsonStatement) FROM(table ReadableTable) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

func (s *selectJsonStatement) WHERE(condition BoolExpression) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

func (s *selectJsonStatement) ORDER_BY(orderBy ...OrderByClause) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

func (s *selectJsonStatement) LIMIT(limit int64) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}

func (s *selectJsonStatement) OFFSET(offset int64) SelectJsonStatement {
	_ = "STUB: not implemented"
	return *new(SelectJsonStatement)
}
