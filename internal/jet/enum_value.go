package jet

// NewEnumValue creates new named enum value
func NewEnumValue(name string) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

type enumValueSerializer struct {
	name string
}

func (e enumValueSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}
