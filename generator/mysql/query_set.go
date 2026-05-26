package mysql

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/metadata"
)

// mySqlQuerySet is dialect query set for MySQL
type mySqlQuerySet struct{}

func (m mySqlQuerySet) GetTablesMetaData(db *sql.DB, schemaName string, tableType metadata.TableType) ([]metadata.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m mySqlQuerySet) GetEnumsMetaData(db *sql.DB, schemaName string) ([]metadata.Enum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
