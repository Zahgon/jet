package postgres

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

// INTERSECT returns all rows that are in query results.
// It eliminates duplicate rows from its result.
func INTERSECT(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

// INTERSECT_ALL returns all rows that are in query results.
// It does not eliminates duplicate rows from its result.
func INTERSECT_ALL(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

// EXCEPT returns all rows that are in the result of query lhs but not in the result of query rhs.
// It eliminates duplicate rows from its result.
func EXCEPT(lhs, rhs jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

// EXCEPT_ALL returns all rows that are in the result of query lhs but not in the result of query rhs.
// It does not eliminates duplicate rows from its result.
func EXCEPT_ALL(lhs, rhs jet.SerializerStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

type setStatement interface {
	setOperators

	ORDER_BY(orderByClauses ...OrderByClause) setStatement

	LIMIT(limit int64) setStatement
	OFFSET(offset int64) setStatement
	// OFFSET_e can be used when an integer expression is needed as offset, otherwise OFFSET can be used
	OFFSET_e(offset IntegerExpression) setStatement

	AsTable(alias string) SelectTable
}

type setOperators interface {
	Statement
	jet.HasProjections
	Expression

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement
	INTERSECT(rhs SelectStatement) setStatement
	INTERSECT_ALL(rhs SelectStatement) setStatement
	EXCEPT(rhs SelectStatement) setStatement
	EXCEPT_ALL(rhs SelectStatement) setStatement
}

type setOperatorsImpl struct {
	stmtRoot setOperators
}

func (s *setOperatorsImpl) UNION(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) UNION_ALL(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) INTERSECT(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) INTERSECT_ALL(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) EXCEPT(rhs SelectStatement) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setOperatorsImpl) EXCEPT_ALL(rhs SelectStatement) setStatement {
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

func (s *setStatementImpl) OFFSET_e(offset IntegerExpression) setStatement {
	_ = "STUB: not implemented"
	return *new(setStatement)
}

func (s *setStatementImpl) AsTable(alias string) SelectTable {
	_ = "STUB: not implemented"
	return *new(SelectTable)
}

const (
	union     = "UNION"
	intersect = "INTERSECT"
	except    = "EXCEPT"
)

func toSelectList(lhs, rhs jet.SerializerStatement, selects ...jet.SerializerStatement) []jet.SerializerStatement {
	_ = "STUB: not implemented"
	return nil
}
