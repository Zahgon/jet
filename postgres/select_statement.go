package postgres

import (
	"math"

	"github.com/go-jet/jet/v2/internal/jet"
)

// RowLock is interface for SELECT statement row lock types
type RowLock = jet.RowLock

// Row lock types
var (
	UPDATE        = jet.NewRowLock("UPDATE")
	NO_KEY_UPDATE = jet.NewRowLock("NO KEY UPDATE")
	SHARE         = jet.NewRowLock("SHARE")
	KEY_SHARE     = jet.NewRowLock("KEY SHARE")
)

// Window function clauses
var (
	PARTITION_BY = jet.PARTITION_BY
	ORDER_BY     = jet.ORDER_BY
	UNBOUNDED    = int64(math.MaxInt64)
	CURRENT_ROW  = jet.CURRENT_ROW
)

// PRECEDING window frame clause
func PRECEDING(offset int64) jet.FrameExtent {
	_ = "STUB: not implemented"
	return *new(jet.FrameExtent)
}

// FOLLOWING window frame clause
func FOLLOWING(offset int64) jet.FrameExtent {
	_ = "STUB: not implemented"
	return *new(jet.FrameExtent)
}

// Window definition reference
var Window = jet.WindowName

// SelectStatement is interface for PostgreSQL SELECT statement
type SelectStatement interface {
	Statement
	jet.HasProjections
	Expression

	DISTINCT(on ...jet.ColumnExpression) SelectStatement
	FROM(tables ...ReadableTable) SelectStatement
	WHERE(expression BoolExpression) SelectStatement
	GROUP_BY(groupByClauses ...GroupByClause) SelectStatement
	HAVING(boolExpression BoolExpression) SelectStatement
	WINDOW(name string) windowExpand
	ORDER_BY(orderByClauses ...OrderByClause) SelectStatement
	LIMIT(limit int64) SelectStatement
	OFFSET(offset int64) SelectStatement
	// OFFSET_e can be used when an integer expression is needed as offset, otherwise OFFSET can be used
	OFFSET_e(offset IntegerExpression) SelectStatement
	FETCH_FIRST(count IntegerExpression) fetchExpand
	FOR(lock RowLock) SelectStatement

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement
	INTERSECT(rhs SelectStatement) setStatement
	INTERSECT_ALL(rhs SelectStatement) setStatement
	EXCEPT(rhs SelectStatement) setStatement
	EXCEPT_ALL(rhs SelectStatement) setStatement

	AsTable(alias string) SelectTable
}

// SELECT creates new SelectStatement with list of projections
func SELECT(projection Projection, projections ...Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func newSelectStatement(stmtType jet.StatementType, table ReadableTable, projections []Projection) *selectStatementImpl {
	_ = "STUB: not implemented"
	return nil
}

type selectStatementImpl struct {
	jet.ExpressionStatement
	setOperatorsImpl

	Select  jet.ClauseSelect
	From    jet.ClauseFrom
	Where   jet.ClauseWhere
	GroupBy jet.ClauseGroupBy
	Having  jet.ClauseHaving
	Window  jet.ClauseWindow
	OrderBy jet.ClauseOrderBy
	Limit   jet.ClauseLimit
	Offset  jet.ClauseOffset
	Fetch   jet.ClauseFetch
	For     jet.ClauseFor
}

func (s *selectStatementImpl) DISTINCT(on ...jet.ColumnExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) FROM(tables ...ReadableTable) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) WHERE(condition BoolExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) GROUP_BY(groupByClauses ...GroupByClause) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) HAVING(boolExpression BoolExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) WINDOW(name string) windowExpand {
	_ = "STUB: not implemented"
	return *new(windowExpand)
}

func (s *selectStatementImpl) ORDER_BY(orderByClauses ...OrderByClause) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) LIMIT(limit int64) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) OFFSET(offset int64) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) OFFSET_e(offset IntegerExpression) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) FETCH_FIRST(count IntegerExpression) fetchExpand {
	_ = "STUB: not implemented"
	return *new(fetchExpand)
}

func (s *selectStatementImpl) FOR(lock RowLock) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) AsTable(alias string) SelectTable {
	_ = "STUB: not implemented"
	return *new(SelectTable)
}

//-----------------------------------------------------

type windowExpand struct {
	selectStatement *selectStatementImpl
	rootStmt        SelectStatement
}

func (w windowExpand) AS(window ...jet.Window) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func toJetFrameOffset(offset int64) jet.Serializer {
	_ = "STUB: not implemented"
	return *new(jet.Serializer)
}

func readableTablesToSerializerList(tables []ReadableTable) []jet.Serializer {
	_ = "STUB: not implemented"
	return nil
}

type fetchExpand struct {
	selectStatement *selectStatementImpl
	rootStmt        SelectStatement
}

func (f fetchExpand) ROWS_ONLY() SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (f fetchExpand) ROWS_WITH_TIES() SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}
