package main

//go:generate sh -c "printf 'package main\n\nconst version = \"'%s'\"\n' $(git describe --tags --abbrev=0) > version.go"

import (
	"flag"
	"fmt"
	"os"
	"slices"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-jet/jet/v2/generator/metadata"
	mysqlgen "github.com/go-jet/jet/v2/generator/mysql"
	postgresgen "github.com/go-jet/jet/v2/generator/postgres"
	sqlitegen "github.com/go-jet/jet/v2/generator/sqlite"
	"github.com/go-jet/jet/v2/generator/template"
	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/utils/errfmt"
	"github.com/go-jet/jet/v2/mysql"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/sqlite"
)

var (
	source string

	dsn        string
	host       string
	port       int
	user       string
	password   string
	sslmode    string
	params     string
	dbName     string
	schemaName string

	ignoreTables string
	ignoreViews  string
	ignoreEnums  string

	skipModel      bool
	skipSQLBuilder bool

	destDir  string
	modelPkg string
	tablePkg string
	viewPkg  string
	enumPkg  string

	tables string
	views  string
	enums  string

	modelJsonTag string
)

type templateFilter struct {
	names  []string
	ignore bool
}

func init() {
	flag.StringVar(&source, "source", "", "Database system name (postgres, mysql, cockroachdb, mariadb or sqlite)")

	flag.StringVar(&dsn, "dsn", "", `Data source name. Unified format for connecting to database.
    	PostgreSQL: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
		Example:
			postgresql://user:pass@localhost:5432/dbname
    	MySQL: https://dev.mysql.com/doc/refman/8.0/en/connecting-using-uri-or-key-value-pairs.html
		Example:
			mysql://jet:jet@tcp(localhost:3306)/dvds
    	SQLite: https://www.sqlite.org/c3ref/open.html#urifilenameexamples
		Example:
			file://path/to/database/file`)
	flag.StringVar(&host, "host", "", "Database host path. Used only if dsn is not set. (Example: localhost)")
	flag.IntVar(&port, "port", 0, "Database port. Used only if dsn is not set.")
	flag.StringVar(&user, "user", "", "Database user. Used only if dsn is not set.")
	flag.StringVar(&password, "password", "", "The user’s password. Used only if dsn is not set.")
	flag.StringVar(&dbName, "dbname", "", "Database name. Used only if dsn is not set.")
	flag.StringVar(&schemaName, "schema", "public", `Database schema name. (default "public")(PostgreSQL only)`)
	flag.StringVar(&params, "params", "", "Additional connection string parameters(optional). Used only if dsn is not set.")
	flag.StringVar(&sslmode, "sslmode", "disable", `Whether or not to use SSL. Used only if dsn is not set. (optional)(default "disable")(PostgreSQL only)`)
	flag.StringVar(&ignoreTables, "ignore-tables", "", `Comma-separated list of tables to ignore.`)
	flag.StringVar(&ignoreViews, "ignore-views", "", `Comma-separated list of views to ignore.`)
	flag.StringVar(&ignoreEnums, "ignore-enums", "", `Comma-separated list of enums to ignore.`)
	flag.BoolVar(&skipModel, "skip-model", false, `Skip model generation.`)
	flag.BoolVar(&skipSQLBuilder, "skip-sql-builder", false, `Skip SQL builder generation.`)

	flag.StringVar(&destDir, "path", "", "Destination directory for files generated.")
	flag.StringVar(&modelPkg, "rel-model-path", "model", "Relative path for the Model files package from the destination directory.")
	flag.StringVar(&tablePkg, "rel-table-path", "table", "Relative path for the Table files package from the destination directory.")
	flag.StringVar(&viewPkg, "rel-view-path", "view", "Relative path for the View files package from the destination directory.")
	flag.StringVar(&enumPkg, "rel-enum-path", "enum", "Relative path for the Enum files package from the destination directory.")
	flag.StringVar(&modelJsonTag, "model-json-tag", "", "Json tag model to be included in Go structs. (optional)(default <empty>)(allowed values: <empty>, pascal-case, camel-case, snake-case")

	flag.StringVar(&tables, "tables", "", `Comma-separated list of tables to generate.`)
	flag.StringVar(&views, "views", "", `Comma-separated list of views to generate.`)
	flag.StringVar(&enums, "enums", "", `Comma-separated list of enums to generate.`)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if dsn == "" && (source == "" || host == "" || port == 0 || user == "" || dbName == "") {
		printErrorAndExit("ERROR: required flag(s) missing")
	}

	if !slices.Contains([]string{"", "snake-case", "pascal-case", "camel-case"}, modelJsonTag) {
		printErrorAndExit("ERROR: json tag does not contain correct value")
	}

	source := getSource()
	tablesFilter := createTemplateFilter(ignoreTables, tables, "tables")
	viewsFilter := createTemplateFilter(ignoreViews, views, "views")
	enumsFilter := createTemplateFilter(ignoreEnums, enums, "enums")

	var err error

	switch source {
	case "postgresql", "postgres", "cockroachdb", "cockroach":
		generatorTemplate := genTemplate(postgres2.Dialect, tablesFilter, viewsFilter, enumsFilter)

		if dsn != "" {
			err = postgresgen.GenerateDSN(dsn, schemaName, destDir, generatorTemplate)
			break
		}

		dbConn := postgresgen.DBConnection{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			SslMode:  sslmode,
			Params:   params,

			DBName:     dbName,
			SchemaName: schemaName,
		}

		err = postgresgen.Generate(
			destDir,
			dbConn,
			generatorTemplate,
		)

	case "mysql", "mysqlx", "mariadb":
		generatorTemplate := genTemplate(mysql.Dialect, tablesFilter, viewsFilter, enumsFilter)

		if dsn != "" {
			err = mysqlgen.GenerateDSN(dsn, destDir, generatorTemplate)
			break
		}

		dbConn := mysqlgen.DBConnection{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Params:   params,
			DBName:   dbName,
		}

		err = mysqlgen.Generate(
			destDir,
			dbConn,
			generatorTemplate,
		)
	case "sqlite":
		if dsn == "" {
			printErrorAndExit("ERROR: required -dsn flag missing.")
		}

		err = sqlitegen.GenerateDSN(
			dsn,
			destDir,
			genTemplate(sqlite.Dialect, tablesFilter, viewsFilter, enumsFilter),
		)

	case "":
		printErrorAndExit("ERROR: required -source or -dsn flag missing.")

	default:
		printErrorAndExit("ERROR: unknown data source " + source + ". Only postgres, mysql, mariadb and sqlite are supported.")
	}

	if err != nil {
		fmt.Println(errfmt.Trace(err))
		os.Exit(2)
	}
}

func usage() { _ = "STUB: not implemented"; return }

func printErrorAndExit(error string) { _ = "STUB: not implemented"; return }

func getSource() string { _ = "STUB: not implemented"; return "" }

func detectSchema(dsn string) string { _ = "STUB: not implemented"; return "" }

// not found

func parseList(list string) []string { _ = "STUB: not implemented"; return nil }

func genTemplate(dialect jet.Dialect, tablesFilter, viewsFilter, enumsFilter templateFilter) template.Template {
	_ = "STUB: not implemented"
	return *new(template.Template)
}

func createTemplateFilter(ignoreList, allowList, filterType string) templateFilter {
	_ = "STUB: not implemented"
	return *new(templateFilter)
}

func shouldSkipTable(table metadata.Table, filter templateFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func shouldSkipEnum(enum metadata.Enum, filter templateFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func createModelTags(columnMetaData metadata.Column) []string {
	_ = "STUB: not implemented"
	return nil
}
