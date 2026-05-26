package jet

// Values hold a set of one or more rows
type Values []RowExpression

func (v Values) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func (v Values) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}
