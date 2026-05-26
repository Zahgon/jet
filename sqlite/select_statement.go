package sqlite

import (
	"github.com/go-jet/jet/v2/internal/jet"
)

// RowLock is interface for SELECT statement row lock types
type RowLock = jet.RowLock

// Row lock types
var (
	UPDATE = jet.NewRowLock("UPDATE")
	SHARE  = jet.NewRowLock("SHARE")
)

// Window function clauses
var (
	PARTITION_BY = jet.PARTITION_BY
	ORDER_BY     = jet.ORDER_BY
	UNBOUNDED    = jet.UNBOUNDED
	CURRENT_ROW  = jet.CURRENT_ROW
)

// PRECEDING window frame clause
func PRECEDING(offset interface{}) jet.FrameExtent {
	_ = "STUB: not implemented"
	return *new(jet.FrameExtent)
}

// FOLLOWING window frame clause
func FOLLOWING(offset interface{}) jet.FrameExtent {
	_ = "STUB: not implemented"
	return *new(jet.FrameExtent)
}

// Window is used to specify window reference from WINDOW clause
var Window = jet.WindowName

// SelectStatement is interface for MySQL SELECT statement
type SelectStatement interface {
	Statement
	jet.HasProjections
	Expression

	DISTINCT() SelectStatement
	FROM(tables ...ReadableTable) SelectStatement
	WHERE(expression BoolExpression) SelectStatement
	GROUP_BY(groupByClauses ...GroupByClause) SelectStatement
	HAVING(boolExpression BoolExpression) SelectStatement
	WINDOW(name string) windowExpand
	ORDER_BY(orderByClauses ...OrderByClause) SelectStatement
	LIMIT(limit int64) SelectStatement
	OFFSET(offset int64) SelectStatement
	FOR(lock RowLock) SelectStatement
	LOCK_IN_SHARE_MODE() SelectStatement

	UNION(rhs SelectStatement) setStatement
	UNION_ALL(rhs SelectStatement) setStatement

	AsTable(alias string) SelectTable
}

// SELECT creates new SelectStatement with list of projections
func SELECT(projection Projection, projections ...Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func newSelectStatement(table ReadableTable, projections []Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

type selectStatementImpl struct {
	jet.ExpressionStatement
	setOperatorsImpl

	Select    jet.ClauseSelect
	From      jet.ClauseFrom
	Where     jet.ClauseWhere
	GroupBy   jet.ClauseGroupBy
	Having    jet.ClauseHaving
	Window    jet.ClauseWindow
	OrderBy   jet.ClauseOrderBy
	Limit     jet.ClauseLimit
	Offset    jet.ClauseOffset
	For       jet.ClauseFor
	ShareLock jet.ClauseOptional
}

func (s *selectStatementImpl) DISTINCT() SelectStatement {
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

func (s *selectStatementImpl) FOR(lock RowLock) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func (s *selectStatementImpl) LOCK_IN_SHARE_MODE() SelectStatement {
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
}

func (w windowExpand) AS(window ...jet.Window) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

func toJetFrameOffset(offset interface{}) jet.Serializer {
	_ = "STUB: not implemented"
	return *new(jet.Serializer)
}

func readableTablesToSerializerList(tables []ReadableTable) []jet.Serializer {
	_ = "STUB: not implemented"
	return nil
}
