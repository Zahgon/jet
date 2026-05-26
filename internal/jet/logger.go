package jet

import (
	"context"
	"time"
)

// PrintableStatement is a statement which sql query can be logged
type PrintableStatement interface {
	Sql() (query string, args []interface{})
	DebugSql() (query string)
}

// LoggerFunc is a function user can implement to support automatic statement logging.
type LoggerFunc func(ctx context.Context, statement PrintableStatement)

var logger LoggerFunc

// SetLoggerFunc sets automatic statement logging
func SetLoggerFunc(loggerFunc LoggerFunc) { _ = "STUB: not implemented"; return }

func callLogger(ctx context.Context, statement Statement) { _ = "STUB: not implemented"; return }

// QueryInfo contains information about executed query
type QueryInfo struct {
	Statement PrintableStatement
	// Depending on how the statement is executed, RowsProcessed is:
	// 	- Number of rows returned for Query() and QueryContext() methods
	// 	- RowsAffected() for Exec() and ExecContext() methods
	// 	- Always 0 for Rows() method.
	RowsProcessed int64
	Duration      time.Duration
	Err           error
}

// QueryLoggerFunc is a function user can implement to retrieve more information about statement executed.
type QueryLoggerFunc func(ctx context.Context, info QueryInfo)

var queryLoggerFunc QueryLoggerFunc

// SetQueryLogger sets automatic query logging function.
func SetQueryLogger(loggerFunc QueryLoggerFunc) { _ = "STUB: not implemented"; return }

func callQueryLoggerFunc(ctx context.Context, info QueryInfo) { _ = "STUB: not implemented"; return }

// Caller returns information about statement caller
func (q QueryInfo) Caller() (file string, line int, function string) {
	_ = "STUB: not implemented"

	// depending on execution type (Query, QueryContext, Exec, ...) looped once or twice
	return "", 0, ""
}
