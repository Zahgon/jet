package postgres

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/template"
)

// DBConnection contains postgres connection details
type DBConnection struct {
	Host string
	Port int
	User string
	// #nosec G117 -- password is used only for the local development
	Password string
	SslMode  string
	Params   string

	DBName     string
	SchemaName string
}

// Generate generates jet files at destination dir from database connection details
func Generate(destDir string, dbConn DBConnection, genTemplate ...template.Template) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GenerateDSN generates jet files using dsn connection string
func GenerateDSN(dsn, schema, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateDB generates jet files using the provided *sql.DB
func GenerateDB(db *sql.DB, schema, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}

func openConnection(dsn string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }
