package sqlite

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/metadata"
)

// sqliteQuerySet is dialect query set for SQLite
type sqliteQuerySet struct{}

func (p sqliteQuerySet) GetTablesMetaData(db *sql.DB, schemaName string, tableType metadata.TableType) ([]metadata.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTableInfoQuery(db *sql.DB) (string, error) { _ = "STUB: not implemented"; return "", nil }

// generated columns were added in version 3.26.0

func (p sqliteQuerySet) GetTableColumnsMetaData(db *sql.DB, schemaName string, tableName string) ([]metadata.Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// stored or virtual column

// will convert VARCHAR(10) -> VARCHAR, etc...
func getColumnType(columnType string) string { _ = "STUB: not implemented"; return "" }

func (p sqliteQuerySet) GetEnumsMetaData(db *sql.DB, schemaName string) ([]metadata.Enum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
