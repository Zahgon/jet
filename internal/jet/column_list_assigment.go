package jet

type expressionOrColumnList interface {
	Serializer
	isExpressionOrColumnList()
}

type columnListAssigment []ColumnAssigment

func (c columnListAssigment) isColumnAssignment() { _ = "STUB: not implemented"; return }

func (c columnListAssigment) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
