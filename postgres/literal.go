package postgres

import (
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
)

// Bool is boolean literal constructor
func Bool(value bool) BoolExpression { _ = "STUB: not implemented"; return *new(BoolExpression) }

// Int is constructor for 64 bit signed integer expressions literals.
var Int = jet.Int

// Int8 is constructor for 8 bit signed integer expressions literals.
func Int8(value int8) IntegerExpression { _ = "STUB: not implemented"; return *new(IntegerExpression) }

// Int16 is constructor for 16 bit signed integer expressions literals.
func Int16(value int16) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Int32 is constructor for 32 bit signed integer expressions literals.
func Int32(value int32) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Int64 is constructor for 64 bit signed integer expressions literals.
func Int64(value int64) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint8 is constructor for 8 bit unsigned integer expressions literals.
func Uint8(value uint8) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint16 is constructor for 16 bit unsigned integer expressions literals.
func Uint16(value uint16) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint32 is constructor for 32 bit unsigned integer expressions literals.
func Uint32(value uint32) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Uint64 is constructor for 64 bit unsigned integer expressions literals.
func Uint64(value uint64) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// Float creates new float literal expression
var Float = jet.Float

// Real is placeholder constructor for 32-bit float literals
func Real(value float32) FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// Double is placeholder constructor for 64-bit float literals
func Double(value float64) FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// Decimal creates new float literal expression
func Decimal(value string) FloatExpression { _ = "STUB: not implemented"; return *new(FloatExpression) }

// String creates new string literal expression
// String is a parameter constructor for the PostgreSQL text type. Using the `Text` constructor is
// generally preferable.
//
// WARNING: String always applies a `text` type cast, which can be problematic if a parameter is compared
// to a `character` column, as this may prevent index usage. In such cases, consider using the Char
// constructor instead. See also other PostgreSQL-specific constructors: Text, Char, and VarChar.
func String(value string) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// Text is a parameter constructor for the PostgreSQL text type. This constructor also adds an
// explicit placeholder type cast to text in the generated query, such as `$3::text`.
// Example usage:
//
//	Text("English")
func Text(value string) StringExpression { _ = "STUB: not implemented"; return *new(StringExpression) }

// Char is a parameter constructor for the PostgreSQL character type. This constructor also adds an
// explicit placeholder type cast to text in the generated query, such as `$3::char(30)`.
// Example usage:
//
//	Char(20)("English")
func Char(length ...int) func(value string) StringExpression { _ = "STUB: not implemented"; return nil }

// VarChar is a parameter constructor for the PostgreSQL character varying type. This constructor
// also adds an explicit placeholder type cast to text in the generated query, such as `$3::varchar(30)`.
// Example usage:
//
//	VarChar(20)("English")
//	VarChar()("English")
func VarChar(length ...int) func(value string) StringExpression {
	_ = "STUB: not implemented"
	return nil
}

// Json creates new json literal expression
func Json(value interface{}) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// UUID is a helper function to create string literal expression from uuid object
// value can be any uuid type with a String method
func UUID(value fmt.Stringer) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// Bytea creates new bytea literal expression
func Bytea(value interface{}) ByteaExpression {
	_ = "STUB: not implemented"
	return *new(ByteaExpression)
}

// Date creates new date literal expression
func Date(year int, month time.Month, day int) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}

// DateT creates new date literal expression from time.Time object
func DateT(t time.Time) DateExpression { _ = "STUB: not implemented"; return *new(DateExpression) }

// Time creates new time literal expression
func Time(hour, minute, second int, nanoseconds ...time.Duration) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

// TimeT creates new time literal expression from time.Time object
func TimeT(t time.Time) TimeExpression { _ = "STUB: not implemented"; return *new(TimeExpression) }

// Timez creates new time with time zone literal expression
func Timez(hour, minute, second int, milliseconds time.Duration, timezone string) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// TimezT creates new time with time zone literal expression from time.Time object
func TimezT(t time.Time) TimezExpression { _ = "STUB: not implemented"; return *new(TimezExpression) }

// Timestamp creates new timestamp literal expression
func Timestamp(year int, month time.Month, day, hour, minute, second int, milliseconds ...time.Duration) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// TimestampT creates new timestamp literal expression from time.Time object
func TimestampT(t time.Time) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// Timestampz creates new timestamp with time zone literal expression
func Timestampz(year int, month time.Month, day, hour, minute, second int, milliseconds time.Duration, timezone string) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// TimestampzT creates new timestamp literal expression from time.Time object
func TimestampzT(t time.Time) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// BoolArray creates new bool array literal expression from list of values
func BoolArray(values ...bool) Array[BoolExpression] { _ = "STUB: not implemented"; return nil }

// Int32Array creates new integer array literal expression from list of values
func Int32Array(values ...int32) Array[IntegerExpression] { _ = "STUB: not implemented"; return nil }

// Int64Array creates new bigint array literal expression from list of values
func Int64Array(values ...int64) Array[IntegerExpression] { _ = "STUB: not implemented"; return nil }

// Float32Array creates new real array literal expression from list of values
func Float32Array(values ...float32) Array[FloatExpression] { _ = "STUB: not implemented"; return nil }

// Float64Array creates new double precision array literal expression from list of values
func Float64Array(values ...float64) Array[FloatExpression] { _ = "STUB: not implemented"; return nil }

// StringArray creates new string array literal expression from list of values
func StringArray(values ...string) Array[StringExpression] { _ = "STUB: not implemented"; return nil }

// ByteaArray creates new bytea array literal expression from list of values
func ByteaArray(values ...[]byte) Array[ByteaExpression] { _ = "STUB: not implemented"; return nil }

// DateArray creates new date array literal expression from list of values
func DateArray(values ...time.Time) Array[DateExpression] { _ = "STUB: not implemented"; return nil }

// TimestampArray creates new timestamp array literal expression from list of values
func TimestampArray(values ...time.Time) Array[TimestampExpression] {
	_ = "STUB: not implemented"
	return nil
}

// TimestampzArray creates new timestampt with timezone array literal expression from list of values
func TimestampzArray(values ...time.Time) Array[TimestampzExpression] {
	_ = "STUB: not implemented"
	return nil
}

// TimeArray creates new time array literal expression from list of values
func TimeArray(values ...time.Time) Array[TimeExpression] { _ = "STUB: not implemented"; return nil }

// TimezArray creates new time with timezone array literal expression from list of values
func TimezArray(values ...time.Time) Array[TimezExpression] { _ = "STUB: not implemented"; return nil }
