package mysql

import (
	"database/sql"

	"github.com/go-jet/jet/v2/generator/template"
)

const mysqlMaxConns = 10

// DBConnection contains MySQL connection details
type DBConnection struct {
	Host string
	Port int
	User string
	// #nosec G117 -- password is used only for the local development
	Password string
	Params   string

	DBName string
}

// Generate generates jet files at destination dir from database connection details
func Generate(destDir string, dbConn DBConnection, generatorTemplate ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateDSN opens connection via DSN string and does everything what Generate does.
func GenerateDSN(dsn, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	// Special case for go mysql driver. It does not understand schema,
	// so we need to trim it before passing to generator
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	return nil
}

func openConnection(connectionString string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateDB generates jet files using the provided *sql.DB
func GenerateDB(db *sql.DB, dbName, destDir string, templates ...template.Template) error {
	_ = "STUB: not implemented"
	return nil
}

// No schemas in MySQL
