package postgres

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/metadata"
)

// postgresQuerySet is dialect query set for PostgreSQL
type postgresQuerySet struct{}

func (p postgresQuerySet) GetTablesMetaData(db *sql.DB, schemaName string, tableType metadata.TableType) ([]metadata.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add materialized views separately, because materialized views are not part of standard information schema

func getColumnsMetaData(db *sql.DB, schemaName string, tableName string) ([]metadata.Column, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p postgresQuerySet) GetEnumsMetaData(db *sql.DB, schemaName string) ([]metadata.Enum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
