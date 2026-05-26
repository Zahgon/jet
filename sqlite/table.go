package sqlite

import "github.com/go-jet/jet/v2/internal/jet"

// Table is interface for MySQL tables
type Table interface {
	jet.SerializerTable
	readableTable

	INSERT(columns ...jet.Column) InsertStatement
	UPDATE(columns ...jet.Column) UpdateStatement
	DELETE() DeleteStatement
}

type readableTable interface {
	// Generates a select query on the current tableName.
	SELECT(projection Projection, projections ...Projection) SelectStatement

	// Creates a inner join tableName Expression using onCondition.
	INNER_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable

	// Creates a left join tableName Expression using onCondition.
	LEFT_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable

	// Creates a right join tableName Expression using onCondition.
	RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable

	// Creates a full join tableName Expression using onCondition.
	FULL_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable

	// Creates a cross join tableName Expression using onCondition.
	CROSS_JOIN(table ReadableTable) joinSelectUpdateTable
}

type joinSelectUpdateTable interface {
	ReadableTable
	UPDATE(columns ...jet.Column) UpdateStatement
}

// ReadableTable interface
type ReadableTable interface {
	readableTable
	jet.Serializer
}

type readableTableInterfaceImpl struct {
	root ReadableTable
}

// Generates a select query on the current tableName.
func (r readableTableInterfaceImpl) SELECT(projection1 Projection, projections ...Projection) SelectStatement {
	_ = "STUB: not implemented"
	return *new(SelectStatement)
}

// Creates a inner join tableName Expression using onCondition.
func (r readableTableInterfaceImpl) INNER_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable {
	_ = "STUB: not implemented"
	return *new(joinSelectUpdateTable)
}

// Creates a left join tableName Expression using onCondition.
func (r readableTableInterfaceImpl) LEFT_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable {
	_ = "STUB: not implemented"
	return *new(joinSelectUpdateTable)
}

// Creates a right join tableName Expression using onCondition.
func (r readableTableInterfaceImpl) RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable {
	_ = "STUB: not implemented"
	return *new(joinSelectUpdateTable)
}

func (r readableTableInterfaceImpl) FULL_JOIN(table ReadableTable, onCondition BoolExpression) joinSelectUpdateTable {
	_ = "STUB: not implemented"
	return *new(joinSelectUpdateTable)
}

func (r readableTableInterfaceImpl) CROSS_JOIN(table ReadableTable) joinSelectUpdateTable {
	_ = "STUB: not implemented"
	return *new(joinSelectUpdateTable)
}

// NewTable creates new table with schema Name, table Name and list of columns
func NewTable(schemaName, name, alias string, columns ...jet.ColumnExpression) Table {
	_ = "STUB: not implemented"
	return *new(Table)
}

type tableImpl struct {
	jet.SerializerTable
	readableTableInterfaceImpl
	root Table
}

func (t *tableImpl) INSERT(columns ...jet.Column) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (t *tableImpl) UPDATE(columns ...jet.Column) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (t *tableImpl) DELETE() DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

type joinTable struct {
	tableImpl
	jet.JoinTable
}

func newJoinTable(lhs jet.Serializer, rhs jet.Serializer, joinType jet.JoinType, onCondition BoolExpression) Table {
	_ = "STUB: not implemented"
	return *new(Table)
}
