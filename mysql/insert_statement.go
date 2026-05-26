package mysql

import "github.com/go-jet/jet/v2/internal/jet"

// InsertStatement is interface for SQL INSERT statements
type InsertStatement interface {
	Statement

	OPTIMIZER_HINTS(hints ...OptimizerHint) InsertStatement

	// Insert row of values
	VALUES(value interface{}, values ...interface{}) InsertStatement
	// Insert row of values, where value for each column is extracted from filed of structure data.
	// If data is not struct or there is no field for every column selected, this method will panic.
	MODEL(data interface{}) InsertStatement
	MODELS(data interface{}) InsertStatement
	AS_NEW() InsertStatement

	ON_DUPLICATE_KEY_UPDATE(assigments ...ColumnAssigment) InsertStatement

	QUERY(selectStatement SelectStatement) InsertStatement

	RETURNING(projections ...Projection) InsertStatement
}

func newInsertStatement(table Table, columns []jet.Column) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

type insertStatementImpl struct {
	jet.SerializerStatement

	Insert         jet.ClauseInsert
	ValuesQuery    jet.ClauseValuesQuery
	Returning      jet.ClauseReturning
	OnDuplicateKey onDuplicateKeyUpdateClause
}

func (is *insertStatementImpl) OPTIMIZER_HINTS(hints ...OptimizerHint) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) VALUES(value interface{}, values ...interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) MODEL(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) MODELS(data interface{}) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (i *insertStatementImpl) RETURNING(projections ...jet.Projection) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) AS_NEW() InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) ON_DUPLICATE_KEY_UPDATE(assigments ...ColumnAssigment) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

func (is *insertStatementImpl) QUERY(selectStatement SelectStatement) InsertStatement {
	_ = "STUB: not implemented"
	return *new(InsertStatement)
}

type onDuplicateKeyUpdateClause []jet.ColumnAssigment

// Serialize for SetClause
func (s onDuplicateKeyUpdateClause) Serialize(statementType jet.StatementType, out *jet.SQLBuilder, options ...jet.SerializeOption) {
	_ = "STUB: not implemented"
	return
}
