package jet

// AND function adds AND operator between expressions. This function can be used, instead of method AND,
// to have a better inlining of a complex condition in the Go code and in the generated SQL.
func AND(expressions ...BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// OR function adds OR operator between expressions. This function can be used, instead of method OR,
// to have a better inlining of a complex condition in the Go code and in the generated SQL.
func OR(expressions ...BoolExpression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// ------------------ Mathematical functions ---------------//

// ABSf calculates absolute value from float expression
func ABSf(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// ABSi calculates absolute value from int expression
func ABSi(integerExpression IntegerExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// POW calculates power of base with exponent
func POW(base, exponent NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// POWER calculates power of base with exponent
func POWER(base, exponent NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// SQRT calculates square root of numeric expression
func SQRT(numericExpression NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// CBRT calculates cube root of numeric expression
func CBRT(numericExpression NumericExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// CEIL calculates ceil of float expression
func CEIL(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// FLOOR calculates floor of float expression
func FLOOR(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// ROUND calculates round of a float expressions with optional precision
func ROUND(floatExpression FloatExpression, precision ...IntegerExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// SIGN returns sign of float expression
func SIGN(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// TRUNC calculates trunc of float expression with optional precision
func TRUNC(floatExpression FloatExpression, precision ...IntegerExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// LN calculates natural algorithm of float expression
func LN(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// LOG calculates logarithm of float expression
func LOG(floatExpression FloatExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// ----------------- Aggregate functions  -------------------//

// AVG is aggregate function used to calculate avg value from numeric expression
func AVG(numericExpression Expression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// BIT_AND is aggregate function used to calculates the bitwise AND of all non-null input values, or null if none.
func BIT_AND(integerExpression IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// BIT_OR is aggregate function used to calculates the bitwise OR of all non-null input values, or null if none.
func BIT_OR(integerExpression IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// BOOL_AND is aggregate function. Returns true if all input values are true, otherwise false
func BOOL_AND(boolExpression BoolExpression) boolWindowExpression {
	_ = "STUB: not implemented"
	return *new(boolWindowExpression)
}

// BOOL_OR is aggregate function. Returns true if at least one input value is true, otherwise false
func BOOL_OR(boolExpression BoolExpression) boolWindowExpression {
	_ = "STUB: not implemented"
	return *new(boolWindowExpression)
}

// COUNT is aggregate function. Returns number of input rows for which the value of expression is not null.
func COUNT(expression Expression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// EVERY is aggregate function. Returns true if all input values are true, otherwise false
func EVERY(boolExpression BoolExpression) boolWindowExpression {
	_ = "STUB: not implemented"
	return *new(boolWindowExpression)
}

// MAX is aggregate function. Returns minimum value of expression across all input values.
func MAX(expression Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// MAXf is aggregate function. Returns maximum value of float expression across all input values
func MAXf(floatExpression FloatExpression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// MAXi is aggregate function. Returns maximum value of int expression across all input values
func MAXi(integerExpression IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// MIN is aggregate function. Returns minimum value of expression across all input values.
func MIN(expression Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// MINf is aggregate function. Returns minimum value of float expression across all input values
func MINf(floatExpression FloatExpression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// MINi is aggregate function. Returns minimum value of int expression across all input values
func MINi(integerExpression IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// SUM is aggregate function. Returns sum of all expressions
func SUM(expression Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// SUMf is aggregate function. Returns sum of expression across all float expressions
func SUMf(floatExpression FloatExpression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// SUMi is aggregate function. Returns sum of expression across all integer expression.
func SUMi(integerExpression IntegerExpression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// ----------------- Window functions  -------------------//

// ROW_NUMBER returns number of the current row within its partition, counting from 1
func ROW_NUMBER() integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// RANK of the current row with gaps; same as row_number of its first peer
func RANK() integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// DENSE_RANK returns rank of the current row without gaps; this function counts peer groups
func DENSE_RANK() integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// PERCENT_RANK calculates relative rank of the current row: (rank - 1) / (total partition rows - 1)
func PERCENT_RANK() floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// CUME_DIST calculates cumulative distribution: (number of partition rows preceding or peer with current row) / total partition rows
func CUME_DIST() floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

// NTILE returns integer ranging from 1 to the argument value, dividing the partition as equally as possible
func NTILE(numOfBuckets int64) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// LAG returns value evaluated at the row that is offset rows before the current row within the partition;
// if there is no such row, instead return default (which must be of the same type as value).
// Both offset and default are evaluated with respect to the current row.
// If omitted, offset defaults to 1 and default to null
func LAG(expr Expression, offsetAndDefault ...interface{}) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

// LEAD returns value evaluated at the row that is offset rows after the current row within the partition;
// if there is no such row, instead return default (which must be of the same type as value).
// Both offset and default are evaluated with respect to the current row.
// If omitted, offset defaults to 1 and default to null
func LEAD(expr Expression, offsetAndDefault ...interface{}) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

// FIRST_VALUE returns value evaluated at the row that is the first row of the window frame
func FIRST_VALUE(value Expression) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

// LAST_VALUE returns value evaluated at the row that is the last row of the window frame
func LAST_VALUE(value Expression) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

// NTH_VALUE returns value evaluated at the row that is the nth row of the window frame (counting from 1); null if no such row
func NTH_VALUE(value Expression, nth int64) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

func leadLagImpl(name string, expr Expression, offsetAndDefault ...interface{}) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

//------------ String functions ------------------//

// HEX function takes an input and returns its equivalent hexadecimal representation
func HEX(expression Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// UNHEX for a string argument str, UNHEX(str) interprets each pair of characters in the argument
// as a hexadecimal number and converts it to the byte represented by the number.
// The return value is a binary string.
func UNHEX(expression StringExpression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// BIT_LENGTH returns number of bits in string expression
func BIT_LENGTH(stringExpression StringOrBlobExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// CHAR_LENGTH returns number of characters in string expression
func CHAR_LENGTH(stringExpression StringOrBlobExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// OCTET_LENGTH returns number of bytes in string expression
func OCTET_LENGTH(stringExpression StringOrBlobExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// LOWER returns string expression in lower case
func LOWER(stringExpression StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// UPPER returns string expression in upper case
func UPPER(stringExpression StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// BTRIM removes the longest string consisting only of characters
// in characters (a space by default) from the start and end of string
func BTRIM(stringExpression StringOrBlobExpression, trimChars ...StringOrBlobExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// LTRIM removes the longest string containing only characters
// from characters (a space by default) from the start of string
func LTRIM(str StringOrBlobExpression, trimChars ...StringOrBlobExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// RTRIM removes the longest string containing only characters
// from characters (a space by default) from the end of string
func RTRIM(str StringOrBlobExpression, trimChars ...StringOrBlobExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// CHR returns character with the given code.
func CHR(integerExpression IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// CONCAT adds two or more expressions together
func CONCAT(expressions ...Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// CONCAT_WS adds two or more expressions together with a separator.
func CONCAT_WS(separator Expression, expressions ...Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// CONVERT converts string to dest_encoding. The original encoding is
// specified by src_encoding. The string must be valid in this encoding.
func CONVERT(str BlobExpression, srcEncoding StringExpression, destEncoding StringExpression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// CONVERT_FROM converts string to the database encoding. The original
// encoding is specified by src_encoding. The string must be valid in this encoding.
func CONVERT_FROM(str BlobExpression, srcEncoding StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// CONVERT_TO converts string to dest_encoding.
func CONVERT_TO(str StringExpression, toEncoding StringExpression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// ENCODE encodes binary data into a textual representation.
// Supported formats are: base64, hex, escape. escape converts zero bytes and
// high-bit-set bytes to octal sequences (\nnn) and doubles backslashes.
func ENCODE(data BlobExpression, format StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// DECODE decodes binary data from textual representation in string.
// Options for format are same as in encode.
func DECODE(data StringExpression, format StringExpression) BlobExpression {
	_ = "STUB: not implemented"
	return *new(BlobExpression)
}

// FORMAT formats a number to a format like "#,###,###.##", rounded to a specified number of decimal places, then it returns the result as a string.
func FORMAT(formatStr StringExpression, formatArgs ...Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// INITCAP converts the first letter of each word to upper case
// and the rest to lower case. Words are sequences of alphanumeric
// characters separated by non-alphanumeric characters.
func INITCAP(str StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// LEFT returns first n characters in the string.
// When n is negative, return all but last |n| characters.
func LEFT(str StringExpression, n IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// RIGHT returns last n characters in the string.
// When n is negative, return all but first |n| characters.
func RIGHT(str StringExpression, n IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// LENGTH returns number of characters in string with a given encoding
func LENGTH(str StringOrBlobExpression, encoding ...StringExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// LPAD fills up the string to length length by prepending the characters
// fill (a space by default). If the string is already longer than length
// then it is truncated (on the right).
func LPAD(str StringExpression, length IntegerExpression, text ...StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// RPAD fills up the string to length length by appending the characters
// fill (a space by default). If the string is already longer than length then it is truncated.
func RPAD(str StringExpression, length IntegerExpression, text ...StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// BIT_COUNT returns the number of bits set in the binary string (also known as “popcount”).
func BIT_COUNT(bytes BlobExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// MD5 calculates the MD5 hash of string, returning the result in hexadecimal
func MD5(stringExpression StringOrBlobExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// REPEAT repeats string the specified number of times
func REPEAT(str StringExpression, n IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// REPLACE replaces all occurrences in string of substring from with substring to
func REPLACE(text, from, to StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// REVERSE returns reversed string.
func REVERSE(stringExpression StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// STRPOS returns location of specified substring (same as position(substring in string),
// but note the reversed argument order)
func STRPOS(str, substring StringExpression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// SUBSTR extracts substring
func SUBSTR(str StringOrBlobExpression, from IntegerExpression, count ...IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// TO_ASCII convert string to ASCII from another encoding
func TO_ASCII(str StringExpression, encoding ...StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// TO_HEX converts number to its equivalent hexadecimal representation
func TO_HEX(number IntegerExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// REGEXP_LIKE Returns 1 if the string expr matches the regular expression specified by the pattern pat, 0 otherwise.
func REGEXP_LIKE(stringExp StringExpression, pattern StringExpression, matchType ...string) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

//----------Range Type Functions ----------------------//

// LOWER_BOUND returns range expressions lower bound. Returns null if range is empty or the requested bound is infinite.
func LOWER_BOUND[T Expression](rangeExpression Range[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// UPPER_BOUND returns range expressions upper bound. Returns null if range is empty or the requested bound is infinite.
func UPPER_BOUND[T Expression](rangeExpression Range[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func rangeTypeCaster[T Expression](rangeExpression Range[T], exp Expression) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// IS_EMPTY returns true if range is empty
func IS_EMPTY[T Expression](rangeExpression Range[T]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// LOWER_INC returns true if lower bound is inclusive. Returns false for empty range.
func LOWER_INC[T Expression](rangeExpression Range[T]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// UPPER_INC returns true if upper bound is inclusive. Returns false for empty range.
func UPPER_INC[T Expression](rangeExpression Range[T]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// LOWER_INF returns true if upper bound is infinite. Returns false for empty range.
func LOWER_INF[T Expression](rangeExpression Range[T]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

// UPPER_INF returns true if lower bound is infinite. Returns false for empty range.
func UPPER_INF[T Expression](rangeExpression Range[T]) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

//----------Data Type Formatting Functions ----------------------//

// TO_CHAR converts expression to string with format
func TO_CHAR(expression Expression, format StringExpression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// TO_DATE converts string to date using format
func TO_DATE(dateStr, format StringExpression) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}

// TO_NUMBER converts string to numeric using format
func TO_NUMBER(floatStr, format StringExpression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// TO_TIMESTAMP converts string to time stamp with time zone using format
func TO_TIMESTAMP(timestampzStr, format StringExpression) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

//----------------- Date/Time Functions and Operators ---------------//

// EXTRACT extracts time component from time expression
func EXTRACT(field string, from Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// CURRENT_DATE returns current date
func CURRENT_DATE() DateExpression { _ = "STUB: not implemented"; return *new(DateExpression) }

// CURRENT_TIME returns current time with time zone
func CURRENT_TIME(precision ...int) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// CURRENT_TIMESTAMP returns current timestamp with time zone
func CURRENT_TIMESTAMP(precision ...int) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// LOCALTIME returns local time of day using optional precision
func LOCALTIME(precision ...int) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

// LOCALTIMESTAMP returns current date and time using optional precision
func LOCALTIMESTAMP(precision ...int) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

// NOW returns current date and time
func NOW() TimestampzExpression { _ = "STUB: not implemented"; return *new(TimestampzExpression) }

// --------------- Conditional Expressions Functions -------------//

// COALESCE function returns the first of its arguments that is not null.
func COALESCE(value Expression, values ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// NULLIF function returns a null value if value1 equals value2; otherwise it returns value1.
func NULLIF(value1, value2 Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// GREATEST selects the largest  value from a list of expressions
func GREATEST(value Expression, values ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

// LEAST selects the smallest  value from a list of expressions
func LEAST(value Expression, values ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

//--------------------------------------------------------------------//

// newFunc creates new function with name and expressions parameters
func newFunc(name string, expressions []Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

type funcSerializer struct {
	name       string
	parameters parametersSerializer
}

func (f *funcSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

func newBoolFunc(name string, expressions ...Expression) BoolExpression {
	_ = "STUB: not implemented"
	return *new(BoolExpression)
}

type parametersSerializer []Expression

func (p parametersSerializer) serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	_ = "STUB: not implemented"
	return
}

// NewFloatWindowFunc creates new float function with name and expressions
func newWindowFunc(name string, expressions ...Expression) windowExpression {
	_ = "STUB: not implemented"
	return *new(windowExpression)
}

// NewFloatWindowFunc creates new float function with name and expressions
func newBoolWindowFunc(name string, expressions ...Expression) boolWindowExpression {
	_ = "STUB: not implemented"
	return *new(boolWindowExpression)
}

// NewFloatFunc creates new float function with name and expressions
func NewFloatFunc(name string, expressions ...Expression) FloatExpression {
	_ = "STUB: not implemented"
	return *new(FloatExpression)
}

// NewFloatWindowFunc creates new float function with name and expressions
func NewFloatWindowFunc(name string, expressions ...Expression) floatWindowExpression {
	_ = "STUB: not implemented"
	return *new(floatWindowExpression)
}

func newIntegerFunc(name string, expressions ...Expression) IntegerExpression {
	_ = "STUB: not implemented"
	return *new(IntegerExpression)
}

// NewFloatWindowFunc creates new float function with name and expressions
func newIntegerWindowFunc(name string, expressions ...Expression) integerWindowExpression {
	_ = "STUB: not implemented"
	return *new(integerWindowExpression)
}

// NewStringFunc creates new string function with name and expression parameters
func NewStringFunc(name string, expressions ...Expression) StringExpression {
	_ = "STUB: not implemented"
	return *new(StringExpression)
}

// NewTimeFunc creates new time function with name and expression parameters
func NewTimeFunc(name string, expressions ...Expression) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

func newTimezFunc(name string, expressions ...Expression) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

// NewTimestampFunc creates new timestamp function with name and expressions
func NewTimestampFunc(name string, expressions ...Expression) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

func newTimestampzFunc(name string, expressions ...Expression) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

// Func can be used to call custom or unsupported database functions.
func Func(name string, expressions ...Expression) Expression {
	_ = "STUB: not implemented"
	return *new(Expression)
}

func NumRange(lowNum, highNum NumericExpression, bounds ...StringExpression) Range[NumericExpression] {
	_ = "STUB: not implemented"
	return nil
}

func Int4Range(lowNum, highNum IntegerExpression, bounds ...StringExpression) Range[Int4Expression] {
	_ = "STUB: not implemented"
	return nil
}

func Int8Range(lowNum, highNum Int8Expression, bounds ...StringExpression) Range[Int8Expression] {
	_ = "STUB: not implemented"
	return nil
}

func TsRange(lowTs, highTs TimestampExpression, bounds ...StringExpression) Range[TimestampExpression] {
	_ = "STUB: not implemented"
	return nil
}

func TstzRange(lowTs, highTs TimestampzExpression, bounds ...StringExpression) Range[TimestampzExpression] {
	_ = "STUB: not implemented"
	return nil
}

func DateRange(lowTs, highTs DateExpression, bounds ...StringExpression) Range[DateExpression] {
	_ = "STUB: not implemented"
	return nil
}

func rangeFuncParamCombiner(low, high Expression, bounds ...StringExpression) []Expression {
	_ = "STUB: not implemented"
	return nil
}

func TimeKeyword(name string) TimeExpression {
	_ = "STUB: not implemented"
	return *new(TimeExpression)
}

func TimezKeyword(name string) TimezExpression {
	_ = "STUB: not implemented"
	return *new(TimezExpression)
}

func TimestampKeyword(name string) TimestampExpression {
	_ = "STUB: not implemented"
	return *new(TimestampExpression)
}

func TimestampzKeyword(name string) TimestampzExpression {
	_ = "STUB: not implemented"
	return *new(TimestampzExpression)
}

func DateKeyword(name string) DateExpression {
	_ = "STUB: not implemented"
	return *new(DateExpression)
}
