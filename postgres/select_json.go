package postgres

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

// SELECT_JSON_ARR creates a new SelectJsonStatement with a list of projections.
func SELECT_JSON_ARR(projections ...Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

// SELECT_JSON_OBJ creates a new SelectJsonStatement with a list of projections.
func SELECT_JSON_OBJ(projections ...Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

type selectJsonStatement struct {
	*selectStatementImpl

	subQuery      *selectStatementImpl
	statementType jet.StatementType
}

func (s *selectJsonStatement) AS(alias string) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (s *selectJsonStatement) FROM(table ...ReadableTable) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) DISTINCT(on ...jet.ColumnExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) WHERE(condition BoolExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) GROUP_BY(groupByClauses ...GroupByClause) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) HAVING(boolExpression BoolExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) WINDOW(name string) windowExpand {
	_ = "STUB: not implemented"
	return *new(windowExpand)
}

func (s *selectJsonStatement) ORDER_BY(orderByClauses ...OrderByClause) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) LIMIT(limit int64) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) OFFSET(offset int64) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) OFFSET_e(offset IntegerExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) FETCH_FIRST(count IntegerExpression) fetchExpand {
	_ = "STUB: not implemented"
	return *new(fetchExpand)
}

func (s *selectJsonStatement) FOR(lock RowLock) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func newSelectStatementJson(projections []Projection, statementType jet.StatementType) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectJsonStatement) setSubQueryAlias(alias string) { _ = "STUB: not implemented"; return }

func constructJsonFunc(statementType jet.StatementType, subQueryAlias string) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}
