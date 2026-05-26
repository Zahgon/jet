package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// Table is interface for MySQL tables
type Table interface {
	readableTable
	writableTable
	jet.SerializerTable
}

type readableTable interface {
	// Generates a select query on the current tableName.
	SELECT(projection Projection, projections ...Projection) SelectStatement

	// Creates a inner join tableName Expression using onCondition.
	INNER_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a left join tableName Expression using onCondition.
	LEFT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a right join tableName Expression using onCondition.
	RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a full join tableName Expression using onCondition.
	FULL_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable

	// Creates a cross join tableName Expression using onCondition.
	CROSS_JOIN(table ReadableTable) ReadableTable
}

type writableTable interface {
	INSERT(columns ...jet.Column) InsertStatement
	UPDATE(columns ...jet.Column) UpdateStatement
	DELETE() DeleteStatement
	LOCK() LockStatement
}

// ReadableTable interface
type ReadableTable interface {
	readableTable
	jet.Serializer
}

// WritableTable interface
type WritableTable interface {
	jet.Table
	writableTable
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
func (r readableTableInterfaceImpl) INNER_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}

// Creates a left join tableName Expression using onCondition.
func (r readableTableInterfaceImpl) LEFT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}

// Creates a right join tableName Expression using onCondition.
func (r readableTableInterfaceImpl) RIGHT_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}

func (r readableTableInterfaceImpl) FULL_JOIN(table ReadableTable, onCondition BoolExpression) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}

func (r readableTableInterfaceImpl) CROSS_JOIN(table ReadableTable) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}

type writableTableInterfaceImpl struct {
	root WritableTable
}

func (w *writableTableInterfaceImpl) INSERT(columns ...jet.Column) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (w *writableTableInterfaceImpl) UPDATE(columns ...jet.Column) UpdateStatement {
	_ = "STUB: not implemented"
	return *new(UpdateStatement)
}

func (w *writableTableInterfaceImpl) DELETE() DeleteStatement {
	_ = "STUB: not implemented"
	return *new(DeleteStatement)
}

func (w *writableTableInterfaceImpl) LOCK() LockStatement {
	_ = "STUB: not implemented"
	return *new(LockStatement)
}

type tableImpl struct {
	readableTableInterfaceImpl
	writableTableInterfaceImpl

	jet.SerializerTable
}

// NewTable creates new table with schema Name, table Name and list of columns
func NewTable(schemaName, name, alias string, columns ...jet.ColumnExpression) Table {
	_ = "STUB: not implemented"
	return *new(Table)
}

type joinTable struct {
	readableTableInterfaceImpl
	jet.JoinTable
}

func newJoinTable(lhs jet.Serializer, rhs jet.Serializer, joinType jet.JoinType, onCondition BoolExpression) ReadableTable {
	_ = "STUB: not implemented"
	return *new(ReadableTable)
}
