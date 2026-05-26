package jet

// SerializerTable interface
type SerializerTable interface {
	Serializer
	Table
}

// Table interface
type Table interface {
	columns() []Column
	SchemaName() string
	TableName() string
	Alias() string
}

// NewTable creates new table with schema Name, table Name and list of columns
func NewTable(schemaName, name, alias string, columns ...ColumnExpression) SerializerTable {
	_ = "STUB: not implemented"
	return *new(SerializerTable)
}

type tableImpl struct {
	schemaName string
	name       string
	alias      string
	columnList []ColumnExpression
}

func (t *tableImpl) SchemaName() string { _ = "STUB: not implemented"; return "" }

func (t *tableImpl) TableName() string { _ = "STUB: not implemented"; return "" }

func (t *tableImpl) columns() []Column { _ = "STUB: not implemented"; return nil }

func (t *tableImpl) Alias() string { _ = "STUB: not implemented"; return "" }

func (t *tableImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Use default schema if the schema name is not set

// JoinType is type of table join
type JoinType int

// Table join types
const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullJoin
	CrossJoin
)

// Join expressions are pseudo readable tables.
type joinTableImpl struct {
	lhs         Serializer
	rhs         Serializer
	joinType    JoinType
	onCondition BoolExpression
}

// JoinTable interface
type JoinTable SerializerTable

// NewJoinTable creates new join table
func NewJoinTable(lhs Serializer, rhs Serializer, joinType JoinType, onCondition BoolExpression) JoinTable {
	_ = "STUB: not implemented"
	return *new(JoinTable)
}

func (t *joinTableImpl) SchemaName() string { _ = "STUB: not implemented"; return "" }

func (t *joinTableImpl) TableName() string { _ = "STUB: not implemented"; return "" }

func (t *joinTableImpl) columns() []Column { _ = "STUB: not implemented"; return nil }

func (t *joinTableImpl) Alias() string { _ = "STUB: not implemented"; return "" }

func (t *joinTableImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
