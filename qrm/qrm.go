package qrm

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"reflect"
)

// Config holds the configuration settings for QRM scanning behavior.
type Config struct {
	// StrictScan, when true, causes the scanning function to panic if it encounters any
	// unused columns in the SQL query result. This ensures that every column is mapped
	// to a field in the destination struct.
	// Does not apply to statements build with SELECT_JSON_OBJ or SELECT_JSON_ARR
	StrictScan bool

	// StrictFieldMapping, when true, causes the scanning function to panic if it encounters any
	// destination struct fields that do not have matching columns in the SQL query result.
	//
	// Optional fields:
	// If a destination field (including struct/slice fields) is not always selected by a query,
	// it can be marked as optional using `qrm:"optional"`. When StrictFieldMapping is enabled,
	// unmapped fields under an optional field will not trigger a panic.
	// Does not apply to statements build with SELECT_JSON_OBJ or SELECT_JSON_ARR
	StrictFieldMapping bool

	// JsonUnmarshalFunc is called by the Query method to unmarshal JSON query results created by
	// SELECT_JSON_OBJ and SELECT_JSON_ARR statements.
	// It can be replaced with any implementation that matches the standard "encoding/json" `Unmarshal` function signature.
	// By default, it uses the `Unmarshal` function from Go's standard `encoding/json` package.
	JsonUnmarshalFunc func(data []byte, v any) error
}

// GlobalConfig is the package-wide configuration for SQL scanning.
// This variable is not thread safe, and it should be modified only once, for instance, during application initialization.
var GlobalConfig = Config{
	StrictScan:         false,
	StrictFieldMapping: false,
	JsonUnmarshalFunc:  json.Unmarshal,
}

// ErrNoRows is returned by Query when query result set is empty
var ErrNoRows = errors.New("qrm: no rows in result set")

// QueryJsonObj executes a SQL query that returns a JSON object, unmarshals the result into the provided destination,
// and returns the number of rows processed.
//
// The query must return exactly one row with a single column; otherwise, an error is returned.
//
// Parameters:
//
//	ctx      - The context for managing query execution (timeouts, cancellations).
//	db       - The database connection or transaction that implements the Queryable interface.
//	query    - The SQL query string to be executed.
//	args     - A slice of arguments to be used with the query.
//	destPtr  - A pointer to the variable where the unmarshaled JSON result will be stored.
//	          The destination should be a pointer to a struct or map[string]any.
//
// Returns:
//
//	rowsProcessed - The number of rows processed by the query execution.
//	err           - An error if query execution or unmarshaling fails.
func QueryJsonObj(ctx context.Context, db Queryable, query string, args []interface{}, destPtr interface{}) (rowsProcessed int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// QueryJsonArr executes a SQL query that returns a JSON array, unmarshals the result into the provided destination,
// and returns the number of rows processed.
//
// The query must return exactly one row with a single column; otherwise, an error is returned.
//
// Parameters:
//
//	ctx      - The context for managing query execution (timeouts, cancellations).
//	db       - The database connection or transaction that implements the Queryable interface.
//	query    - The SQL query string to be executed.
//	args     - A slice of arguments to be used with the query.
//	destPtr  - A pointer to the variable where the unmarshaled JSON array will be stored.
//	          The destination should be a pointer to a slice of structs or []map[string]any.
//
// Returns:
//
//	rowsProcessed - The number of rows processed by the query execution.
//	err           - An error if query execution or unmarshaling fails.
func QueryJsonArr(ctx context.Context, db Queryable, query string, args []interface{}, destPtr interface{}) (rowsProcessed int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var jsonDestObjErr = "jet: SELECT_JSON_OBJ destination has to be a pointer to struct or pointer to map[string]any"
var jsonDestArrErr = "jet: SELECT_JSON_ARR destination has to be a pointer to slice of struct or pointer to []map[string]any"

func queryJson(ctx context.Context, db Queryable, query string, args []interface{}, destPtr interface{}) (rowsProcessed int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Query executes a Query Result Mapping (QRM) of the provided SQL `query` with a list of parameterized arguments `args`
// over the database connection `db` using the provided context `ctx` and stores the result in the destination `destPtr`.
//
// The destination must be a pointer to either a struct or a slice of structs
// If the destination is a pointer to a struct and no rows are returned, the method returns qrm.ErrNoRows.
//
// Parameters:
//
//	ctx      - The context for managing query execution (timeouts, cancellations).
//	db       - The database connection or transaction implementing the Queryable interface.
//	query    - The SQL query string to be executed.
//	args     - A slice of arguments to be used with the query.
//	destPtr  - A pointer to the variable where the query result will be stored. This can be a pointer to a struct or a slice of structs.
//
// Returns:
//
//	rowsProcessed - The number of rows processed by the query execution.
//	err           - An error if query execution or result mapping fails, or if no rows are found when a struct is expected.
func Query(ctx context.Context, db Queryable, query string, args []interface{}, destPtr interface{}) (rowsProcessed int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// edge case when row result set contains only NULLs.

// ScanOneRowToDest will scan one row into struct destination
func ScanOneRowToDest(scanContext *ScanContext, rows *sql.Rows, destPtr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func queryToSlice(ctx context.Context, db Queryable, query string, args []interface{}, slicePtr interface{}) (rowsProcessed int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func mapRowToSlice(
	scanContext *ScanContext,
	groupKey string,
	slicePtrValue reflect.Value,
	field *reflect.StructField) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mapRowToBaseTypeSlice(scanContext *ScanContext, slicePtrValue reflect.Value, field *reflect.StructField) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mapRowToStruct(
	scanContext *ScanContext,
	groupKey string,
	structPtrValue reflect.Value,
	parentField *reflect.StructField,
	onlySlices ...bool, // small optimization, not to assign to already assigned struct fields
) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// private field

// scannedValue is nil, destination should be set to zero value

// simple type

func qrmAssignError(scannedValue reflect.Value, field reflect.StructField, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func mapRowToDestinationValue(
	scanContext *ScanContext,
	groupKey string,
	dest reflect.Value,
	structField *reflect.StructField) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mapRowToDestinationPtr(
	scanContext *ScanContext,
	groupKey string,
	destPtrValue reflect.Value,
	structField *reflect.StructField) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
