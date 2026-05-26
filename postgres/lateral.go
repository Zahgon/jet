package postgres

// LATERAL derived tables constructor from select statement
func LATERAL(selectStmt SelectStatement) lateralImpl {
	_ = "STUB: not implemented"
	return *new(lateralImpl)
}

type lateralImpl struct {
	selectStmt SelectStatement
}

func (l lateralImpl) AS(alias string) SelectTable {
	_ = "STUB: not implemented"
	return *new(SelectTable)
}
