// Modeling of columns

package jet

// Column is common column interface for all types of columns.
type Column interface {
	Name() string
	TableName() string

	setTableName(table string)
	setSubQuery(subQuery SelectTable)
	defaultAlias() string
}

// ColumnSerializer is interface for all serializable columns
type ColumnSerializer interface {
	Serializer
	Column
}

// ColumnExpression interface
type ColumnExpression interface {
	Column
	Expression
}

// ColumnExpressionImpl is base type for sql columns.
type ColumnExpressionImpl struct {
	ExpressionInterfaceImpl

	name      string
	tableName string

	subQuery SelectTable
}

// NewColumnImpl creates new ColumnExpressionImpl
func NewColumnImpl(name string, tableName string, root ColumnExpression) *ColumnExpressionImpl {
	_ = "STUB: not implemented"
	return nil
}

// Name returns name of the column
func (c *ColumnExpressionImpl) Name() string {
	_ = "STUB: not implemented"

	// TableName returns column table name
	return ""
}

func (c *ColumnExpressionImpl) TableName() string { _ = "STUB: not implemented"; return "" }

func (c *ColumnExpressionImpl) setTableName(table string) { _ = "STUB: not implemented"; return }

func (c *ColumnExpressionImpl) setSubQuery(subQuery SelectTable) { _ = "STUB: not implemented"; return }

func (c *ColumnExpressionImpl) defaultAlias() string { _ = "STUB: not implemented"; return "" }

func (c *ColumnExpressionImpl) serializeForOrderBy(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// set Statement (UNION, EXCEPT ...) can reference only select projections in order by clause
//always quote

func (c *ColumnExpressionImpl) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (c *ColumnExpressionImpl) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (c *ColumnExpressionImpl) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (c *ColumnExpressionImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
