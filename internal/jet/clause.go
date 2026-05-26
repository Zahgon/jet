package jet

// Clause interface
type Clause interface {
	Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption)
}

// ClauseWithProjections interface
type ClauseWithProjections interface {
	Clause

	Projections() ProjectionList
}

// OptimizerHint provides a way to optimize query execution per-statement basis
type OptimizerHint string

type optimizerHints []OptimizerHint

func (o optimizerHints) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseSelect struct
type ClauseSelect struct {
	Distinct          bool
	DistinctOnColumns []ColumnExpression
	ProjectionList    []Projection

	IsForRowToJson bool

	// MySQL only
	OptimizerHints optimizerHints
}

// Projections returns list of projections for select clause
func (s *ClauseSelect) Projections() ProjectionList {
	_ = "STUB: not implemented"
	return *

	// Serialize serializes clause into SQLBuilder
	new(ProjectionList)
}

func (s *ClauseSelect) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseFrom struct
type ClauseFrom struct {
	Name   string
	Tables []Serializer
}

// Serialize serializes clause into SQLBuilder
func (f *ClauseFrom) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
	// SELECT statement does not have to have FROM clause
}

// ClauseWhere struct
type ClauseWhere struct {
	Condition BoolExpression
	Mandatory bool
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseWhere) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseGroupBy struct
type ClauseGroupBy struct {
	List []GroupByClause
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseGroupBy) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseHaving struct
type ClauseHaving struct {
	Condition BoolExpression
}

// Serialize serializes clause into SQLBuilder
func (c *ClauseHaving) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseOrderBy struct
type ClauseOrderBy struct {
	List        []OrderByClause
	SkipNewLine bool
}

func (o *ClauseOrderBy) serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Serialize serializes clause into SQLBuilder
func (o *ClauseOrderBy) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseLimit struct
type ClauseLimit struct {
	Count int64
}

func (o *ClauseLimit) serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Serialize serializes clause into SQLBuilder
func (l *ClauseLimit) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseOffset struct
type ClauseOffset struct {
	Count IntegerExpression
}

func (o *ClauseOffset) serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Serialize serializes clause into SQLBuilder
func (o *ClauseOffset) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseFetch struct
type ClauseFetch struct {
	Count    IntegerExpression
	WithTies bool
}

// Serialize serializes ClauseFetch into sql builder output
func (o *ClauseFetch) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseFor struct
type ClauseFor struct {
	Lock RowLock
}

// Serialize serializes clause into SQLBuilder
func (f *ClauseFor) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseSetStmtOperator struct
type ClauseSetStmtOperator struct {
	Operator       string
	All            bool
	Selects        []SerializerStatement
	OrderBy        ClauseOrderBy
	Limit          ClauseLimit
	Offset         ClauseOffset
	SkipSelectWrap bool
}

// Projections returns set of projections for ClauseSetStmtOperator
func (s *ClauseSetStmtOperator) Projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

// Serialize serializes clause into SQLBuilder
func (s *ClauseSetStmtOperator) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseUpdate struct
type ClauseUpdate struct {
	Table SerializerTable

	// MySQL only
	OptimizerHints optimizerHints
}

// Serialize serializes clause into SQLBuilder
func (u *ClauseUpdate) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// SetClause struct
type SetClause struct {
	Columns []Column
	Values  []Serializer
}

// Serialize serializes clause into SQLBuilder
func (s *SetClause) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseInsert struct
type ClauseInsert struct {
	Table   SerializerTable
	Columns []Column

	// MySQL only
	OptimizerHints optimizerHints
}

// GetColumns gets list of columns for insert
func (i *ClauseInsert) GetColumns() []Column { _ = "STUB: not implemented"; return nil }

// Serialize serializes clause into SQLBuilder
func (i *ClauseInsert) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseValuesQuery struct
type ClauseValuesQuery struct {
	ClauseValues
	ClauseQuery
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseValuesQuery) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseValues struct
type ClauseValues struct {
	Rows [][]Serializer
	As   string
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseValues) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseQuery struct
type ClauseQuery struct {
	Query          SerializerStatement
	SkipSelectWrap bool
}

// Serialize serializes clause into SQLBuilder
func (v *ClauseQuery) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseDelete struct
type ClauseDelete struct {
	Table SerializerTable

	// MySQL only
	OptimizerHints optimizerHints
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseDelete) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseStatementBegin struct
type ClauseStatementBegin struct {
	Name   string
	Tables []SerializerTable
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseStatementBegin) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseOptional struct
type ClauseOptional struct {
	Name      string
	Show      bool
	InNewLine bool
}

// Serialize serializes clause into SQLBuilder
func (d *ClauseOptional) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseIn struct
type ClauseIn struct {
	LockMode string
}

// Serialize serializes clause into SQLBuilder
func (i *ClauseIn) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// WindowDefinition struct
type WindowDefinition struct {
	Name   string
	Window Window
}

// ClauseWindow struct
type ClauseWindow struct {
	Definitions []WindowDefinition
}

// Serialize serializes clause into SQLBuilder
func (i *ClauseWindow) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// SetPair clause
type SetPair struct {
	Column ColumnSerializer
	Value  Serializer
}

// SetClauseNew clause
type SetClauseNew []ColumnAssigment

// Serialize for SetClauseNew
func (s SetClauseNew) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// KeywordClause type
type KeywordClause struct {
	Keyword
}

// Serialize for KeywordClause
func (k KeywordClause) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// ClauseReturning  type
type ClauseReturning struct {
	ProjectionList []Projection
}

// Serialize for ClauseReturning
func (r *ClauseReturning) Serialize(statementType StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Projections for ClauseReturning
func (r ClauseReturning) Projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}
