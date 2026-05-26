package sqlite

// CAST function converts an expr (of any type) into later specified datatype.
func CAST(expr Expression) *cast { _ = "STUB: not implemented"; return nil }

type cast struct {
	expr Expression
}

// AS casts expressions to castType
func (c *cast) AS(castType string) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// AS_TEXT cast expression to TEXT type
func (c *cast) AS_TEXT() StringExpression { _ = "STUB: not implemented"; return *new(StringExpression) }

// AS_NUMERIC cast expression to NUMERIC type
func (c *cast) AS_NUMERIC() FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// AS_INTEGER cast expression to INTEGER type
func (c *cast) AS_INTEGER() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_REAL cast expression to REAL type
func (c *cast) AS_REAL() FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// AS_BLOB cast expression to BLOB type
func (c *cast) AS_BLOB() BlobExpression { _ = "STUB: not implemented"; return *new(BlobExpression) }
