package jet

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/qrm"
)

// Statement is a common interface for all SQL statements, including SELECT, SELECT_JSON_ARR, SELECT_JSON_OBJ, INSERT,
// UPDATE, DELETE, and LOCK.
type Statement interface {
	// Sql returns a parameterized SQL query along with its list of arguments.
	Sql() (query string, args []interface{})

	// DebugSql returns a debug-friendly SQL query where all parameterized placeholders
	// are replaced with their respective argument string representations.
	//
	// Warning: This method should only be used for debugging purposes.
	//   Do not use it in production, as it may lead to security risks such as SQL injection.
	DebugSql() (query string)

	// Query delegates call to QueryContext using context.Background() as parameter.
	Query(db qrm.Queryable, destination interface{}) error

	// QueryContext executes the statement with the provided context over a database connection or transaction (`db`),
	// and stores the retrieved row results in the given destination.
	//
	// For statements of type SELECT, INSERT, UPDATE, or DELETE, the destination must be a pointer to either a struct or a slice.
	// For SELECT_JSON_ARR statements, the destination must be a pointer to a slice of structs or a pointer to []map[string]any.
	// For SELECT_JSON_OBJ statements, the destination must be a pointer to a struct or a pointer to map[string]any.
	//
	// If the destination is a pointer to a struct and the query returns no rows, QueryContext returns qrm.ErrNoRows.
	QueryContext(ctx context.Context, db qrm.Queryable, destination interface{}) error

	// Exec delegates call to ExecContext using context.Background() as parameter.
	Exec(db qrm.Executable) (sql.Result, error)

	// ExecContext executes statement with context over db connection/transaction without returning any rows.
	ExecContext(ctx context.Context, db qrm.Executable) (sql.Result, error)

	// Rows executes statements over db connection/transaction and returns rows
	Rows(ctx context.Context, db qrm.Queryable) (*Rows, error)
}

// Rows wraps sql.Rows type with a support for query result mapping
type Rows struct {
	*sql.Rows

	scanContext *qrm.ScanContext
}

// Scan will map the Row values into struct destination
func (r *Rows) Scan(destination interface{}) error { _ = "STUB: not implemented"; return nil }

// SerializerStatement interface
type SerializerStatement interface {
	Serializer
	Statement
	HasProjections
}

// HasProjections interface
type HasProjections interface {
	projections() ProjectionList
}

// SerializerHasProjections interface is combination of Serializer and HasProjections interface
type SerializerHasProjections interface {
	Serializer
	HasProjections
}

// statementInterfaceImpl struct
type statementInterfaceImpl struct {
	dialect       Dialect
	statementType StatementType
	root          SerializerStatement
}

func (s *statementInterfaceImpl) Sql() (query string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *statementInterfaceImpl) DebugSql() (query string) { _ = "STUB: not implemented"; return "" }

func (s *statementInterfaceImpl) Query(db qrm.Queryable, destination interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *statementInterfaceImpl) QueryContext(ctx context.Context, db qrm.Queryable, destination interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *statementInterfaceImpl) query(
	ctx context.Context,
	queryFunc func(query string, args []interface{}) (int64, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *statementInterfaceImpl) Exec(db qrm.Executable) (res sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s *statementInterfaceImpl) ExecContext(ctx context.Context, db qrm.Executable) (res sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s *statementInterfaceImpl) Rows(ctx context.Context, db qrm.Queryable) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func duration(f func()) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// ExpressionStatement interfacess
type ExpressionStatement interface {
	Expression
	Statement
	HasProjections
}

// NewExpressionStatementImpl creates new expression statement
func NewExpressionStatementImpl(Dialect Dialect,
	statementType StatementType,
	root ExpressionStatement,
	clauses ...Clause) ExpressionStatement {
	_ = "STUB: not implemented"
	return *new(ExpressionStatement)
}

type expressionStatementImpl struct {
	ExpressionInterfaceImpl
	statementImpl
}

func (s *expressionStatementImpl) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (e *expressionStatementImpl) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// NewStatementImpl creates new statementImpl
func NewStatementImpl(Dialect Dialect, statementType StatementType, root SerializerStatement, clauses ...Clause) SerializerStatement {
	_ = "STUB: not implemented"
	return *new(SerializerStatement)
}

type statementImpl struct {
	statementInterfaceImpl

	Clauses []Clause
}

func (s *statementImpl) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

func (s *statementImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
