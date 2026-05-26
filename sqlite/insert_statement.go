package sqlite

import "github.com/go-jet/jet/v2/internal/jet"

// InsertStatement is interface for SQL INSERT statements
type InsertStatement interface {
	Statement

	VALUES(value interface{}, values ...interface{}) InsertStatement
	MODEL(data interface{}) InsertStatement
	MODELS(data interface{}) InsertStatement
	QUERY(selectStatement SelectStatement) InsertStatement
	DEFAULT_VALUES() InsertStatement

	ON_CONFLICT(indexExpressions ...jet.ColumnExpression) onConflict
	RETURNING(projections ...Projection) InsertStatement
}

func newInsertStatement(table Table, columns []jet.Column) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

type insertStatementImpl struct {
	jet.SerializerStatement

	Insert        jet.ClauseInsert
	ValuesQuery   jet.ClauseValuesQuery
	DefaultValues jet.ClauseOptional
	OnConflict    onConflictClause
	Returning     jet.ClauseReturning
}

func (is *insertStatementImpl) VALUES(value interface{}, values ...interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

// MODEL will insert row of values, where value for each column is extracted from filed of structure data.
// If data is not struct or there is no field for every column selected, this method will panic.
func (is *insertStatementImpl) MODEL(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) MODELS(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) QUERY(selectStatement SelectStatement) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) DEFAULT_VALUES() InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) RETURNING(projections ...jet.Projection) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) ON_CONFLICT(indexExpressions ...jet.ColumnExpression) onConflict {
	_ = "STUB: not implemented"
	return *new(onConflict)
}
