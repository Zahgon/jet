package jet

import (
	"fmt"
	"time"
)

type literalSerializer struct {
	value    interface{}
	constant bool
}

func (l *literalSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Literal is injected directly to SQL query, and does not appear in parametrized argument list.
func Literal(value interface{}) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// FixedLiteral is injected directly to SQL query, and does not appear in parametrized argument list.
func FixedLiteral(value interface{}) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// Int creates a new 64 bit signed integer literal
func Int(value int64) IntegerExpression { _ = "STUB: not implemented"; return *new(IntegerExpression) }

// Int8 creates a new 8 bit signed integer literal
func Int8(value int8) IntegerExpression { _ = "STUB: not implemented"; return *new(IntegerExpression) }

// Int16 creates a new 16 bit signed integer literal
func Int16(value int16) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Int32 creates a new 32 bit signed integer literal
func Int32(value int32) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint8 creates a new 8 bit unsigned integer literal
func Uint8(value uint8) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint16 creates a new 16 bit unsigned integer literal
func Uint16(value uint16) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint32 creates a new 32 bit unsigned integer literal
func Uint32(value uint32) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint64 creates a new 64 bit unsigned integer literal
func Uint64(value uint64) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Bool creates new bool literal expression
func Bool(value bool) BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// Float creates new float literal from float64 value
func Float(value float64) FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// Decimal creates new float literal from string value
func Decimal(value string) FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// String creates new string literal expression
func String(value string) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// Time creates new time literal expression
func Time(hour, minute, second int, nanoseconds ...time.Duration) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

// TimeT creates new time literal expression from time.Time object
func TimeT(t time.Time) TimeExpression { _ = "STUB: not implemented"; return *new(TimeExpression) }

// Timez creates new time with time zone literal expression
func Timez(hour, minute, second int, nanoseconds time.Duration, timezone string) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// TimezT creates new time with time zone literal expression from time.Time object
func TimezT(t time.Time) TimezExpression { _ = "STUB: not implemented"; return *new(TimezExpression) }

// Timestamp creates new timestamp literal expression
func Timestamp(year int, month time.Month, day, hour, minute, second int, nanoseconds ...time.Duration) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// TimestampT creates new timestamp literal expression from time.Time object
func TimestampT(t time.Time) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// Timestampz creates new timestamp with time zone literal expression
func Timestampz(year int, month time.Month, day, hour, minute, second int, nanoseconds time.Duration, timezone string) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// TimestampzT creates new timestamp literal expression from time.Time object
func TimestampzT(t time.Time) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// Date creates new date literal expression
func Date(year int, month time.Month, day int) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}

// DateT creates new date literal expression from time.Time object
func DateT(t time.Time) DateExpression { _ = "STUB: not implemented"; return *new(DateExpression) }

func formatNanoseconds(nanoseconds ...time.Duration) string { _ = "STUB: not implemented"; return "" }

//--------------------------------------------------//

var (
	// NULL is jet equivalent of SQL NULL
	NULL = newExpression(Keyword("NULL"))
	// STAR is jet equivalent of SQL *
	STAR = newExpression(Keyword("*"))
	// PLUS_INFINITY is jet equivalent for sql infinity
	PLUS_INFINITY = String("infinity")
	// MINUS_INFINITY is jet equivalent for sql -infinity
	MINUS_INFINITY = String("-infinity")
)

//---------------------------------------------------//

type rawSerializer struct {
	Raw           string
	NamedArgument map[string]interface{}
}

func (n *rawSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// Raw can be used for any unsupported functions, operators or expressions.
// For example: Raw("current_database()")
func Raw(raw string, namedArgs ...map[string]interface{}) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// RawBool helper that for raw string boolean expressions
func RawBool(raw string, namedArgs ...map[string]interface{}) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// RawInt helper that for integer expressions
func RawInt(raw string, namedArgs ...map[string]interface{}) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// RawFloat helper that for float expressions
func RawFloat(raw string, namedArgs ...map[string]interface{}) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// RawString helper that for string expressions
func RawString(raw string, namedArgs ...map[string]interface{}) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// RawTime helper that for time expressions
func RawTime(raw string, namedArgs ...map[string]interface{}) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

// RawTimez helper that for time with time zone expressions
func RawTimez(raw string, namedArgs ...map[string]interface{}) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// RawTimestamp helper that for timestamp expressions
func RawTimestamp(raw string, namedArgs ...map[string]interface{}) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// RawTimestampz helper that for timestamp with time zone expressions
func RawTimestampz(raw string, namedArgs ...map[string]interface{}) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// RawDate helper that for date expressions
func RawDate(raw string, namedArgs ...map[string]interface{}) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}

// RawBlob is raw query helper that for blob expressions
func RawBlob(raw string, namedArgs ...map[string]interface{}) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// RawRange helper that for range expressions
func RawRange[T Expression](raw string, namedArgs ...map[string]interface{}) Range[T] {
	_ = "STUB: not implemented"
	return nil
}

// UUID is a helper function to create string literal expression from uuid object
// value can be any uuid type with a String method
func UUID(value fmt.Stringer) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}
