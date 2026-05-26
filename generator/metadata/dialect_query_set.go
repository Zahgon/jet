package metadata

import (
	"database/sql"
)

// TableType is type of database table(view or base)
type TableType string

// SQL table types
const (
	BaseTable TableType = "BASE TABLE"
	ViewTable TableType = "VIEW"
)

// DialectQuerySet is set of methods necessary to retrieve dialect metadata information
type DialectQuerySet interface {
	GetTablesMetaData(db *sql.DB, schemaName string, tableType TableType) ([]Table, error)
	GetEnumsMetaData(db *sql.DB, schemaName string) ([]Enum, error)
}

// GetSchema retrieves Schema information from database
func GetSchema(db *sql.DB, querySet DialectQuerySet, schemaName string) (Schema, error) {
	_ = "STUB: not implemented"
	return *new(Schema), nil
}
