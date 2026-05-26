package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-jet/jet/v2/internal/utils/errfmt"

	"github.com/go-jet/jet/v2/tests/dbconfig"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var testSuite string

func init() {
	flag.StringVar(&testSuite, "testsuite", "all", "Test suite name (postgres, mysql, mariadb, cockroach, sqlite or all)")
	flag.Parse()
}

// Database names
const (
	Postgres  = "postgres"
	MySql     = "mysql"
	MariaDB   = "mariadb"
	Sqlite    = "sqlite"
	Cockroach = "cockroach"
)

func main() {
	var err error

	switch strings.ToLower(testSuite) {
	case Postgres:
		err = initPostgresDB(Postgres, dbconfig.PostgresConnectString)
	case Cockroach:
		err = initPostgresDB(Cockroach, dbconfig.CockroachConnectString)
	case MySql:
		err = initMySQLDB(false)
	case MariaDB:
		err = initMySQLDB(true)
	case Sqlite:
		err = initSQLiteDB()
	case "all":
		err = initPostgresDB(Cockroach, dbconfig.CockroachConnectString)
		if err != nil {
			break
		}
		err = initPostgresDB(Postgres, dbconfig.PostgresConnectString)
		if err != nil {
			break
		}
		err = initMySQLDB(false)
		if err != nil {
			break
		}

		err = initMySQLDB(true)
		if err != nil {
			break
		}
		err = initSQLiteDB()
	default:
		panic("invalid testsuite flag. Test suite name (postgres, mysql, mariadb, cockroach, sqlite or all)")
	}

	if err != nil {
		fmt.Println(errfmt.Trace(err))
		os.Exit(1)
	}
}

func initSQLiteDB() error { _ = "STUB: not implemented"; return nil }

func initMySQLDB(isMariaDB bool) error { _ = "STUB: not implemented"; return nil }

// #nosec G204

func initPostgresDB(dbType string, connectionString string) error {
	_ = "STUB: not implemented"
	return nil
}

// retry add due to a concurrency issue in CockroachDB, specifically a TransactionRetryError

func retry(count int, f func() error) error { _ = "STUB: not implemented"; return nil }

func execFile(db *sql.DB, sqlFilePath string) error { _ = "STUB: not implemented"; return nil }

// #nosec G304

func execInTx(db *sql.DB, f func(tx *sql.Tx) error) error { _ = "STUB: not implemented"; return nil }

// to speed up initialization of test database
