package stmtcache

import (
	"context"
	"database/sql"
)

// Tx is a wrapper around *sql.Tx, adding prepared statement caching capability.
// Tx is not thread safe and should not be shared between goroutines.
type Tx struct {
	*sql.Tx

	db         *DB
	statements map[string]*sql.Stmt
}

// Exec executes a query that doesn't return rows. Exec delegates call to ExecContext with contex.Background()
// as parameter.
func (t *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// ExecContext executes a query that doesn't return rows. If statement caching is enabled, ExecContext will
// first call PrepareContext to retrieve a prepared statement, and then execute a query using a prepared statement.
// If statement caching is disabled, this method delegates the call to the *sql.Tx ExecContext method.
func (t *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query delegates call to QueryContext using context.Background() as parameter.
func (t *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryContext executes a query that returns rows. If statement caching is enabled, QueryContext will
// first call PrepareContext to retrieve a prepared statement, and then execute a query using a prepared statement.
// If statement caching is disabled, this method delegates the call to the *sql.Tx QueryContext method.
func (t *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare delegates call to PrepareContext using context.Background as a parameter.
func (t *Tx) Prepare(query string) (*sql.Stmt, error) { _ = "STUB: not implemented"; return nil, nil }

// PrepareContext returns database prepared statement for a query. When statement caching is enabled, it returns a cached
// prepared statement if available; otherwise, it creates a new prepared statement and adds it to the cache.
// Invoking this method directly is unnecessary, as wrapper methods like Exec/ExecContext and Query/QueryContext
// will call PrepareContext before executing a query on it.
// If statement caching is disabled, this method delegates the call to the *sql.Tx PrepareContext method.
//
// There's no need to manually close the returned statement; it operates within the transaction scope and will be closed
// automatically upon the completion of the transaction, whether it's committed or rolled back.
func (t *Tx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
