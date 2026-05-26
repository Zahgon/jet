package jet

// SelectTable is interface for SELECT sub-queries
type SelectTable interface {
	SerializerHasProjections
	Alias() string
	AllColumns() ProjectionList
}

type selectTableImpl struct {
	Statement     SerializerHasProjections
	alias         string
	columnAliases []ColumnExpression
}

// NewSelectTable func
func NewSelectTable(selectStmt SerializerHasProjections, alias string, columnAliases []ColumnExpression) selectTableImpl {
	_ = "STUB: not implemented"
	return *new(selectTableImpl)
}

func (s selectTableImpl) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

func (s selectTableImpl) Alias() string { _ = "STUB: not implemented"; return "" }

func (s selectTableImpl) AllColumns() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

func (s selectTableImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// --------------------------------------

type lateralImpl struct {
	selectTableImpl
}

// NewLateral creates new lateral expression from select statement with alias
func NewLateral(selectStmt SerializerStatement, alias string) SelectTable {
	_ = "STUB: not implemented"
	return *new(SelectTable)
}

func (s lateralImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
