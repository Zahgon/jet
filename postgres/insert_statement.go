package postgres

import "github.com/go-jet/jet/v2/internal/jet"

// InsertStatement is interface for SQL INSERT statements
type InsertStatement interface {
	jet.SerializerStatement

	// Insert row of values
	VALUES(value interface{}, values ...interface{}) InsertStatement
	// Insert row of values, where value for each column is extracted from filed of structure data.
	// If data is not struct or there is no field for every column selected, this method will panic.
	MODEL(data interface{}) InsertStatement
	MODELS(data interface{}) InsertStatement
	QUERY(selectStatement SelectStatement) InsertStatement

	ON_CONFLICT(indexExpressions ...jet.ColumnExpression) onConflict

	RETURNING(projections ...Projection) InsertStatement
}

func newInsertStatement(table WritableTable, columns []jet.Column) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

type insertStatementImpl struct {
	jet.SerializerStatement

	Insert      jet.ClauseInsert
	ValuesQuery jet.ClauseValuesQuery
	Returning   jet.ClauseReturning
	OnConflict  onConflictClause
}

func (i *insertStatementImpl) VALUES(value interface{}, values ...interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) MODEL(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) MODELS(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) RETURNING(projections ...jet.Projection) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) QUERY(selectStatement SelectStatement) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) ON_CONFLICT(indexExpressions ...jet.ColumnExpression) onConflict {
	_ = "STUB: not implemented"
	return *new(onConflict)
}
