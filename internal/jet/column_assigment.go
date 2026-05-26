package jet

// ColumnAssigment is interface wrapper around column assignment
type ColumnAssigment interface {
	Serializer
	isColumnAssignment()
}

type columnAssigmentImpl struct {
	column   ColumnSerializer
	toAssign Serializer
}

func (a columnAssigmentImpl) isColumnAssignment() { _ = "STUB: not implemented"; return }

func (a columnAssigmentImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
