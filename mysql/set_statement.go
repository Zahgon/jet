package mysql

import "github.com/go-jet/jet/v2/internal/jet"

// UNION effectively appends the result of sub-queries(select statements) into single query.
// It eliminates duplicate rows from its result.
func UNION(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

// UNION_ALL effectively appends the result of sub-queries(select statements) into single query.
// It does not eliminates duplicate rows from its result.
func UNION_ALL(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

type setStatement interface {
	setOperators

	ORDER_BY(orderByClauses ...OrderByClause) setStatement

	LIMIT(limit int64) setStatement
	OFFSET(offset int64) setStatement

	AsTable(alias string) SelectTable
}

type setOperators interface {
	jet.Statement
	jet.HasProjections
	jet.Expression

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement
}

type setOperatorsImpl struct {
	root setOperators
}

func (s *setOperatorsImpl) UNION(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) UNION_ALL(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

type setStatementImpl struct {
	jet.ExpressionStatement

	setOperatorsImpl

	setOperator jet.ClauseSetStmtOperator
}

func newSetStatementImpl(operator string, all bool, selects []jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setStatementImpl) ORDER_BY(orderByClauses ...OrderByClause) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setStatementImpl) LIMIT(limit int64) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setStatementImpl) OFFSET(offset int64) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setStatementImpl) AsTable(alias string) SelectTable {
	_ = "STUB: not implemented"
	return *new(SelectTable)
}

const (
	union = "UNION"
)

func toSelectList(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) []jet.SerializerStatement {
	_ = "STUB: not implemented"
	return nil
}
