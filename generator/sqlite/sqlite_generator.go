package sqlite

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/template"
)

// GenerateDSN generates jet files using dsn connection string
func GenerateDSN(dsn, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateDB generates jet files using the provided *sql.DB
func GenerateDB(db *sql.DB, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}
