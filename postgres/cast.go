package postgres

// CAST function converts an expr (of any type) into later specified datatype.
func CAST(expr Expression) *cast { _ = "STUB: not implemented"; return nil }

type cast struct {
	expr Expression
}

func (b *cast) AS(castType string) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// AS_BOOL casts expression as bool type
func (b *cast) AS_BOOL() BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// AS_SMALLINT casts expression as smallint type
func (b *cast) AS_SMALLINT() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_INTEGER casts expression AS integer type
func (b *cast) AS_INTEGER() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_BIGINT casts expression AS bigint type
func (b *cast) AS_BIGINT() IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// AS_NUMERIC casts expression as numeric type, using precision and optionally scale
func (b *cast) AS_NUMERIC(precisionAndScale ...int) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// AS_REAL casts expression AS real type
func (b *cast) AS_REAL() FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// AS_DOUBLE casts expression AS double precision type
func (b *cast) AS_DOUBLE() FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// AS_TEXT casts expression AS text type
func (b *cast) AS_TEXT() StringExpression { _ = "STUB: not implemented"; return *new(StringExpression) }

// AS_CHAR casts expression AS a character type
func (b *cast) AS_CHAR(length ...int) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// AS_VARCHAR casts expression AS a character varying type
func (b *cast) AS_VARCHAR(length ...int) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// AS_DATE casts expression AS date type
func (b *cast) AS_DATE() DateExpression { _ = "STUB: not implemented"; return *new(DateExpression) }

// AS_DECIMAL casts expression AS decimal type
func (b *cast) AS_DECIMAL() FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// AS_BYTEA casts expression AS bytea type
func (b *cast) AS_BYTEA() ByteaExpression { _ = "STUB: not implemented"; return *new(ByteaExpression) }

// AS_TIME casts expression AS date type
func (b *cast) AS_TIME() TimeExpression { _ = "STUB: not implemented"; return *new(TimeExpression) }

// AS_TIMEZ casts expression AS time with time timezone type
func (b *cast) AS_TIMEZ() TimezExpression { _ = "STUB: not implemented"; return *new(TimezExpression) }

// AS_TIMESTAMP casts expression AS timestamp type
func (b *cast) AS_TIMESTAMP() TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// AS_TIMESTAMPZ casts expression AS timestamp with timezone type
func (b *cast) AS_TIMESTAMPZ() TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// AS_INTERVAL casts expression AS interval type
func (b *cast) AS_INTERVAL() IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

// AS_UUID casts expression AS uuid type
func (b *cast) AS_UUID() StringExpression { _ = "STUB: not implemented"; return *new(StringExpression) }

// AS_BOOL_ARRAY casts expression as boolean array type
func (b *cast) AS_BOOL_ARRAY() Array[BoolExpression] { _ = "STUB: not implemented"; return nil }

// AS_INTEGER_ARRAY casts expression as integer array type
func (b *cast) AS_INTEGER_ARRAY() Array[IntegerExpression] { _ = "STUB: not implemented"; return nil }

// AS_BIGINT_ARRAY casts expression as bigint array type
func (b *cast) AS_BIGINT_ARRAY() Array[IntegerExpression] { _ = "STUB: not implemented"; return nil }

// AS_REAL_ARRAY casts expression as real array
func (b *cast) AS_REAL_ARRAY() Array[FloatExpression] { _ = "STUB: not implemented"; return nil }

// AS_DOUBLE_ARRAY casts expression as double precision array
func (b *cast) AS_DOUBLE_ARRAY() Array[FloatExpression] { _ = "STUB: not implemented"; return nil }

// AS_TEXT_ARRAY casts expression as text array
func (b *cast) AS_TEXT_ARRAY() Array[StringExpression] { _ = "STUB: not implemented"; return nil }

// AS_BYTEA_ARRAY casts expression as bytea array
func (b *cast) AS_BYTEA_ARRAY() Array[ByteaExpression] { _ = "STUB: not implemented"; return nil }

// AS_DATE_ARRAY casts expression as date array
func (b *cast) AS_DATE_ARRAY() Array[DateExpression] { _ = "STUB: not implemented"; return nil }

// AS_TIMESTAMP_ARRAY casts expression as timestamp array
func (b *cast) AS_TIMESTAMP_ARRAY() Array[TimestampExpression] {
	_ = "STUB: not implemented"
	return nil
}

// AS_TIMESTAMPZ_ARRAY casts expression as timestamp with time zone array
func (b *cast) AS_TIMESTAMPZ_ARRAY() Array[TimestampzExpression] {
	_ = "STUB: not implemented"
	return nil
}

// AS_TIME_ARRAY casts expression as time array
func (b *cast) AS_TIME_ARRAY() Array[TimeExpression] { _ = "STUB: not implemented"; return nil }

// AS_TIMEZ_ARRAY casts expression as time with timezone array
func (b *cast) AS_TIMEZ_ARRAY() Array[TimezExpression] { _ = "STUB: not implemented"; return nil }
