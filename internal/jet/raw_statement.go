package jet

type rawStatementImpl struct {
	statementInterfaceImpl

	RawQuery       string
	NamedArguments map[string]interface{}
}

// RawStatement creates new sql statements from raw query and optional map of named arguments
func RawStatement(dialect Dialect, rawQuery string, namedArgument ...map[string]interface{}) SerializerStatement {
	_ = "STUB: not implemented"
	return *new(SerializerStatement)
}

func (s *rawStatementImpl) projections() ProjectionList {
	_ = "STUB: not implemented"
	return *new(ProjectionList)
}

func (s *rawStatementImpl) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
