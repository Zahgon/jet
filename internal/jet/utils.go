package jet

// SerializeClauseList func
func SerializeClauseList(statement StatementType, clauses []Serializer, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

func serializeExpressionList(
	statement StatementType,
	expressions []Expression,
	separator string,
	out *SQLBuilder,
	options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// SerializeProjectionList func
func SerializeProjectionList(statement StatementType, projections []Projection, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// SerializeProjectionListJsonObj serializes a list of projections for JSON object
func SerializeProjectionListJsonObj(statement StatementType, projections []Projection, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// SerializeColumnNames func
func SerializeColumnNames(columns []Column, out *SQLBuilder) { _ = "STUB: not implemented"; return }

// SerializeColumnExpressions func
func SerializeColumnExpressions(columns []ColumnExpression, statementType StatementType,
	out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// SerializeColumnExpressionNames func
func SerializeColumnExpressionNames(columns []ColumnExpression, out *SQLBuilder) {
	_ = "STUB: not implemented"
	return
}

// ToSerializerList converts list of expressions to list of serializers
func ToSerializerList[T Serializer](elems []T) []Serializer { _ = "STUB: not implemented"; return nil }

// ToExpressionList converts list of any expressions to list of expressions
func ToExpressionList[T Expression](expressions []T) []Expression {
	_ = "STUB: not implemented"
	return nil
}

// ColumnListToProjectionList func
func ColumnListToProjectionList(columns []ColumnExpression) []Projection {
	_ = "STUB: not implemented"
	return nil
}

// ToSerializerValue creates Serializer type from the value
func ToSerializerValue(value interface{}) Serializer {
	_ = "STUB: not implemented"
	return *new(Serializer)
}

// UnwindRowFromModel func
func UnwindRowFromModel(columns []Column, data interface{}) []Serializer {
	_ = "STUB: not implemented"
	return nil
}

// UnwindRowsFromModels func
func UnwindRowsFromModels(columns []Column, data interface{}) [][]Serializer {
	_ = "STUB: not implemented"
	return nil
}

// UnwindRowFromValues func
func UnwindRowFromValues(value interface{}, values []interface{}) []Serializer {
	_ = "STUB: not implemented"
	return nil
}

// UnwidColumnList func
func UnwidColumnList(columns []Column) []Column { _ = "STUB: not implemented"; return nil }

// OptionalOrDefaultString will return first value from variable argument list str or
// defaultStr if variable argument list is empty
func OptionalOrDefaultString(defaultStr string, str ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// OptionalOrDefault will return first value from variable argument list expression or
// defaultExpression if variable argument list is empty
func OptionalOrDefault(expressions []Expression, defaultExpression Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func extractTableAndColumnName(alias string) (tableName string, columnName string) {
	_ = "STUB: not implemented"
	return "", ""
}

func serializeToDefaultDebugString(expr Serializer) string { _ = "STUB: not implemented"; return "" }

// joinAlias examples:
//
//	joinAlias("foo", "bar") // "foo.bar"
//	joinAlias("foo.*", "bar") // "foo.bar"
//	joinAlias("", "bar") // "bar"
func joinAlias(tableAlias, columnAlias string) string { _ = "STUB: not implemented"; return "" }

func singleOptional[T any](value []T) T { _ = "STUB: not implemented"; return *new(T) }
