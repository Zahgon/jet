package mysql

// CAST function converts an expr (of any type) into later specified datatype.
func CAST(expr Expression) *cast { _ = "STUB: not implemented"; return nil }

type cast struct {
	expr Expression
}

// AS casts expressions to castType
func (c *cast) AS(castType string) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// AS_DATETIME cast expression to DATETIME type
func (c *cast) AS_DATETIME() DateTimeExpression {
	_ = "STUB: not implemented"
	return *new(DateTimeExpression)
}

// AS_SIGNED casts expression to SIGNED type
func (c *cast) AS_SIGNED() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_UNSIGNED casts expression to UNSIGNED type
func (c *cast) AS_UNSIGNED() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_CHAR casts expression to CHAR type with optional length
func (c *cast) AS_CHAR(length ...int) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// AS_DATE casts expression AS DATE type
func (c *cast) AS_DATE() DateExpression { _ = "STUB: not implemented"; return *new(DateExpression) }

func (c *cast) AS_FLOAT() FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

func (c *cast) AS_DOUBLE() FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// AS_DECIMAL casts expression AS DECIMAL type
func (c *cast) AS_DECIMAL() FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// AS_TIME casts expression AS TIME type
func (c *cast) AS_TIME() TimeExpression { _ = "STUB: not implemented"; return *new(TimeExpression) }

// AS_BINARY casts expression as BINARY type
func (c *cast) AS_BINARY() BlobExpression { _ = "STUB: not implemented"; return *new(BlobExpression) }
