package testutils

import (
	"context"
	"testing"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-jet/jet/v2/stmtcache"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// UnixTimeComparer will compare time equality while ignoring time zone
var UnixTimeComparer = cmp.Comparer(func(t1, t2 time.Time) bool {
	return t1.Unix() == t2.Unix()
})

// AssertExecAndRollback will execute and rollback statement in sql transaction
func AssertExecAndRollback(t *testing.T, stmt jet.Statement, db *stmtcache.DB, rowsAffected ...int64) {
	_ = "STUB: not implemented"
	return
}

// AssertExec assert statement execution for successful execution and number of rows affected
func AssertExec(t *testing.T, stmt jet.Statement, db qrm.DB, rowsAffected ...int64) {
	_ = "STUB: not implemented"
	return
}

// AssertExecContext assert statement execution for successful execution and number of rows affected
func AssertExecContext(t *testing.T, stmt jet.Statement, ctx context.Context, db qrm.DB, rowsAffected ...int64) {
	_ = "STUB: not implemented"
	return
}

// ExecuteInTxAndRollback will execute function in sql transaction and then rollback transaction
func ExecuteInTxAndRollback(t *testing.T, db *stmtcache.DB, f func(tx qrm.DB)) {
	_ = "STUB: not implemented"
	return
}

// AssertExecErr assert statement execution for failed execution with error string errorStr
func AssertExecErr(t *testing.T, stmt jet.Statement, db qrm.DB, errorStr string) {
	_ = "STUB: not implemented"
	return
}

// AssertExecContextErr assert statement execution for failed execution with error string errorStr
func AssertExecContextErr(ctx context.Context, t *testing.T, stmt jet.Statement, db qrm.DB, errorStr string) {
	_ = "STUB: not implemented"
	return
}

func getFullPath(relativePath string) string { _ = "STUB: not implemented"; return "" }

// PrintJson print v as json
func PrintJson(v interface{}) { _ = "STUB: not implemented"; return }

// ToJSON converts v into json string
func ToJSON(v interface{}) string { _ = "STUB: not implemented"; return "" }

// AssertJSON check if data json output is the same as expectedJSON
func AssertJSON(t *testing.T, data interface{}, expectedJSON string) {
	_ = "STUB: not implemented"
	return
}

// AssertJsonEqual checks if actual and expected json representation are the same
func AssertJsonEqual(t require.TestingT, actual, expected interface{}, option ...cmp.Option) {
	_ = "STUB: not implemented"
	return
}

// SaveJSONFile saves v as json at testRelativePath
// nolint:unused
func SaveJSONFile(v interface{}, testRelativePath string) { _ = "STUB: not implemented"; return }

func ReadJSONFile(t require.TestingT, testRelativePath string, dest any) {
	_ = "STUB: not implemented"
	return
}

// skip assert for benchmarks

// #nosec G304

// AssertJSONFile check if data json representation is the same as json at testRelativePath
func AssertJSONFile(t require.TestingT, data interface{}, testRelativePath string) {
	_ = "STUB: not implemented"
	return
}

// skip assert for benchmarks

// #nosec G304

//AssertDeepEqual(t, string(fileJSONData), string(jsonData))

// AssertStatementSql check if statement Sql() is the same as expectedQuery and expectedArgs
func AssertStatementSql(t require.TestingT, query jet.PrintableStatement, expectedQuery string, expectedArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// skip assert for benchmarks

// AssertStatementSqlErr checks if statement Sql() panics with errorStr
func AssertStatementSqlErr(t *testing.T, stmt jet.Statement, errorStr string) {
	_ = "STUB: not implemented"
	return
}

// AssertDebugStatementSql check if statement Sql() is the same as expectedQuery
func AssertDebugStatementSql(t *testing.T, query jet.PrintableStatement, expectedQuery string, expectedArgs ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// AssertSerialize checks if clause serialize produces expected query and args
func AssertSerialize(t *testing.T, dialect jet.Dialect, serializer jet.Serializer, query string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

//fmt.Println(out.Buff.String())

// AssertDebugSerialize checks if clause serialize produces expected debug query and args
func AssertDebugSerialize(t *testing.T, dialect jet.Dialect, clause jet.Serializer, query string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// AssertClauseSerialize checks if clause serialize produces expected query and args
func AssertClauseSerialize(t *testing.T, dialect jet.Dialect, clause jet.Clause, query string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// AssertPanicErr checks if running a function fun produces a panic with errorStr string
func AssertPanicErr(t *testing.T, fun func(), errorStr string) { _ = "STUB: not implemented"; return }

// AssertSerializeErr check if clause serialize panics with errString
func AssertSerializeErr(t *testing.T, dialect jet.Dialect, clause jet.Serializer, errString string) {
	_ = "STUB: not implemented"
	return
}

// AssertProjectionSerialize check if projection serialize produces expected query and args
func AssertProjectionSerialize(t *testing.T, dialect jet.Dialect, projection jet.Projection, query string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// AssertQueryPanicErr check if statement Query execution panics with error errString
func AssertQueryPanicErr(t *testing.T, stmt jet.Statement, db qrm.DB, dest interface{}, errString string) {
	_ = "STUB: not implemented"
	return
}

// AssertFileContent check if file content at filePath contains expectedContent text.
func AssertFileContent(t *testing.T, filePath string, expectedContent string) {
	_ = "STUB: not implemented"
	return
}

// #nosec G304

// AssertFileNamesEqual check if all filesInfos are contained in fileNames
func AssertFileNamesEqual(t *testing.T, dirPath string, fileNames ...string) {
	_ = "STUB: not implemented"
	return
}

// DeepCopy create deep copy of src
func DeepCopy[T any](t require.TestingT, src T) T { _ = "STUB: not implemented"; return *new(T) }

// AssertDeepEqual checks if actual and expected objects are deeply equal.
func AssertDeepEqual(t require.TestingT, actual, expected interface{}, option ...cmp.Option) {
	_ = "STUB: not implemented"
	return
}

func assertQueryString(t require.TestingT, actual, expected string) {
	_ = "STUB: not implemented"
	return
}

func printDiff(actual, expected interface{}, options ...cmp.Option) {
	_ = "STUB: not implemented"
	return
}

// UUIDPtr returns address of uuid.UUID
func UUIDPtr(u string) *uuid.UUID { _ = "STUB: not implemented"; return nil }
