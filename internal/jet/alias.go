package jet

type alias struct {
	expression Expression
	alias      string
}

func newAlias(expression Expression, aliasName string) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (a *alias) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	// if alias is in the form "table.column", we break it into two parts so that ProjectionList.As(newAlias) can
	// overwrite tableName with a new alias. This method is called only for exporting aliased custom columns.
	// Generated columns have default aliasing.
	return *new(Projection)
}

// This function is used to create dummy columns when exporting sub-query columns using subQuery.AllColumns()
// In most case we don't care about type of the column, except when sub-query columns are used as SELECT_JSON projection.
// We need to know type to encode value for json unmarshal.
func newDummyColumnForExpression(exp Expression, name string) ColumnExpression {
	_ = "STUB: not implemented"
	return *new(ColumnExpression)
}

func (a *alias) serializeForProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (a *alias) serializeForJsonObjEntry(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func (a *alias) serializeForRowToJsonProjection(statement StatementType, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}
