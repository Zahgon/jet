package jet

// ColumnBool is interface for SQL boolean columns.
type ColumnBool interface {
	BoolExpression
	Column

	From(subQuery SelectTable) ColumnBool
	SET(boolExp BoolExpression) ColumnAssigment
}

type boolColumnImpl struct {
	boolInterfaceImpl
	*ColumnExpressionImpl
}

func (i *boolColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *boolColumnImpl) From(subQuery SelectTable) ColumnBool {
	_ = "STUB: not implemented"
	return *new(ColumnBool)
}

func (i *boolColumnImpl) SET(boolExp BoolExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// BoolColumn creates named bool column.
func BoolColumn(name string) ColumnBool { _ = "STUB: not implemented"; return *new(ColumnBool) }

//------------------------------------------------------//

// ColumnFloat is interface for SQL real, numeric, decimal or double precision column.
type ColumnFloat interface {
	FloatExpression
	Column

	From(subQuery SelectTable) ColumnFloat
	SET(floatExp FloatExpression) ColumnAssigment
}

type floatColumnImpl struct {
	floatInterfaceImpl
	*ColumnExpressionImpl
}

func (i *floatColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *floatColumnImpl) From(subQuery SelectTable) ColumnFloat {
	_ = "STUB: not implemented"
	return *new(ColumnFloat)
}

func (i *floatColumnImpl) SET(floatExp FloatExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// FloatColumn creates named float column.
func FloatColumn(name string) ColumnFloat { _ = "STUB: not implemented"; return *new(ColumnFloat) }

//------------------------------------------------------//

// ColumnInteger is interface for SQL smallint, integer, bigint columns.
type ColumnInteger interface {
	IntegerExpression
	Column

	From(subQuery SelectTable) ColumnInteger
	SET(intExp IntegerExpression) ColumnAssigment
}

type integerColumnImpl struct {
	integerInterfaceImpl

	*ColumnExpressionImpl
}

func (i *integerColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *integerColumnImpl) From(subQuery SelectTable) ColumnInteger {
	_ = "STUB: not implemented"
	return *new(ColumnInteger)
}

func (i *integerColumnImpl) SET(intExp IntegerExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// IntegerColumn creates named integer column.
func IntegerColumn(name string) ColumnInteger {
	_ = "STUB: not implemented"
	return *new(ColumnInteger)
}

//------------------------------------------------------//

type ColumnArray[E Expression] interface {
	Array[E]
	Column

	From(subQuery SelectTable) ColumnArray[E]
	SET(stringExp Array[E]) ColumnAssigment
}

type arrayColumnImpl[E Expression] struct {
	arrayInterfaceImpl[E]

	*ColumnExpressionImpl
}

func (a arrayColumnImpl[E]) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (a arrayColumnImpl[E]) From(subQuery SelectTable) ColumnArray[E] {
	_ = "STUB: not implemented"
	return nil
}

func (a *arrayColumnImpl[E]) SET(stringExp Array[E]) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// StringColumn creates named string column.
func ArrayColumn[E Expression](name string) ColumnArray[E] { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------//

// ColumnString is interface for SQL text, character, character varying
// uuid columns and enums types.
type ColumnString interface {
	StringExpression
	Column

	From(subQuery SelectTable) ColumnString
	SET(stringExp StringExpression) ColumnAssigment
}

type stringColumnImpl struct {
	stringInterfaceImpl

	*ColumnExpressionImpl
}

func (i *stringColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *stringColumnImpl) From(subQuery SelectTable) ColumnString {
	_ = "STUB: not implemented"
	return *new(ColumnString)
}

func (i *stringColumnImpl) SET(stringExp StringExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// StringColumn creates named string column.
func StringColumn(name string) ColumnString { _ = "STUB: not implemented"; return *new(ColumnString) }

//------------------------------------------------------//

// ColumnBlob is interface for binary data types (bytea, binary, blob, etc...)
type ColumnBlob interface {
	BlobExpression
	Column

	From(subQuery SelectTable) ColumnBlob
	SET(blob BlobExpression) ColumnAssigment
}

type blobColumnImpl struct {
	blobInterfaceImpl

	*ColumnExpressionImpl
}

func (i *blobColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *blobColumnImpl) From(subQuery SelectTable) ColumnBlob {
	_ = "STUB: not implemented"
	return *new(ColumnBlob)
}

func (i *blobColumnImpl) SET(blobExp BlobExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// BlobColumn creates named blob column.
func BlobColumn(name string) ColumnBlob { _ = "STUB: not implemented"; return *new(ColumnBlob) }

//------------------------------------------------------//

// ColumnTime is interface for SQL time column.
type ColumnTime interface {
	TimeExpression
	Column

	From(subQuery SelectTable) ColumnTime
	SET(timeExp TimeExpression) ColumnAssigment
}

type timeColumnImpl struct {
	timeInterfaceImpl
	*ColumnExpressionImpl
}

func (i *timeColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *timeColumnImpl) From(subQuery SelectTable) ColumnTime {
	_ = "STUB: not implemented"
	return *new(ColumnTime)
}

func (i *timeColumnImpl) SET(timeExp TimeExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// TimeColumn creates named time column
func TimeColumn(name string) ColumnTime { _ = "STUB: not implemented"; return *new(ColumnTime) }

//------------------------------------------------------//

// ColumnTimez is interface of SQL time with time zone columns.
type ColumnTimez interface {
	TimezExpression
	Column

	From(subQuery SelectTable) ColumnTimez
	SET(timeExp TimezExpression) ColumnAssigment
}

type timezColumnImpl struct {
	timezInterfaceImpl
	*ColumnExpressionImpl
}

func (i *timezColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *timezColumnImpl) From(subQuery SelectTable) ColumnTimez {
	_ = "STUB: not implemented"
	return *new(ColumnTimez)
}

func (i *timezColumnImpl) SET(timezExp TimezExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// TimezColumn creates named time with time zone column.
func TimezColumn(name string) ColumnTimez { _ = "STUB: not implemented"; return *new(ColumnTimez) }

//------------------------------------------------------//

// ColumnTimestamp is interface of SQL timestamp columns.
type ColumnTimestamp interface {
	TimestampExpression
	Column

	From(subQuery SelectTable) ColumnTimestamp
	SET(timestampExp TimestampExpression) ColumnAssigment
}

type timestampColumnImpl struct {
	timestampInterfaceImpl
	*ColumnExpressionImpl
}

func (i *timestampColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *timestampColumnImpl) From(subQuery SelectTable) ColumnTimestamp {
	_ = "STUB: not implemented"
	return *new(ColumnTimestamp)
}

func (i *timestampColumnImpl) SET(timestampExp TimestampExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// TimestampColumn creates named timestamp column
func TimestampColumn(name string) ColumnTimestamp {
	_ = "STUB: not implemented"
	return *new(ColumnTimestamp)
}

//------------------------------------------------------//

// ColumnTimestampz is interface of SQL timestamp with timezone columns.
type ColumnTimestampz interface {
	TimestampzExpression
	Column

	From(subQuery SelectTable) ColumnTimestampz
	SET(timestampzExp TimestampzExpression) ColumnAssigment
}

type timestampzColumnImpl struct {
	timestampzInterfaceImpl
	*ColumnExpressionImpl
}

func (i *timestampzColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *timestampzColumnImpl) From(subQuery SelectTable) ColumnTimestampz {
	_ = "STUB: not implemented"
	return *new(ColumnTimestampz)
}

func (i *timestampzColumnImpl) SET(timestampzExp TimestampzExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// TimestampzColumn creates named timestamp with time zone column.
func TimestampzColumn(name string) ColumnTimestampz {
	_ = "STUB: not implemented"
	return *new(ColumnTimestampz)
}

//------------------------------------------------------//

// ColumnDate is interface of SQL date columns.
type ColumnDate interface {
	DateExpression
	Column

	From(subQuery SelectTable) ColumnDate
	SET(dateExp DateExpression) ColumnAssigment
}

type dateColumnImpl struct {
	dateInterfaceImpl
	*ColumnExpressionImpl
}

func (i *dateColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *dateColumnImpl) From(subQuery SelectTable) ColumnDate {
	_ = "STUB: not implemented"
	return *new(ColumnDate)
}

func (i *dateColumnImpl) SET(dateExp DateExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// DateColumn creates named date column.
func DateColumn(name string) ColumnDate { _ = "STUB: not implemented"; return *new(ColumnDate) }

//------------------------------------------------------//

// ColumnInterval is interface of PostgreSQL interval columns.
type ColumnInterval interface {
	IntervalExpression
	Column

	From(subQuery SelectTable) ColumnInterval
	SET(intervalExp IntervalExpression) ColumnAssigment
}

//------------------------------------------------------//

type intervalColumnImpl struct {
	*ColumnExpressionImpl
	intervalInterfaceImpl
}

func (i *intervalColumnImpl) SET(intervalExp IntervalExpression) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

func (i *intervalColumnImpl) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *intervalColumnImpl) From(subQuery SelectTable) ColumnInterval {
	_ = "STUB: not implemented"
	return *new(ColumnInterval)
}

// IntervalColumn creates named interval column.
func IntervalColumn(name string) ColumnInterval {
	_ = "STUB: not implemented"
	return *new(ColumnInterval)
}

//------------------------------------------------------//

// ColumnRange is interface for range columns which can be int range, string range
// timestamp range or date range.
type ColumnRange[T Expression] interface {
	Range[T]
	Column

	From(subQuery SelectTable) ColumnRange[T]
	SET(rangeExp Range[T]) ColumnAssigment
}

type rangeColumnImpl[T Expression] struct {
	rangeInterfaceImpl[T]
	*ColumnExpressionImpl
}

func (i *rangeColumnImpl[T]) fromImpl(subQuery SelectTable) Projection {
	_ = "STUB: not implemented"
	return *new(Projection)
}

func (i *rangeColumnImpl[T]) From(subQuery SelectTable) ColumnRange[T] {
	_ = "STUB: not implemented"
	return nil
}

func (i *rangeColumnImpl[T]) SET(rangeExp Range[T]) ColumnAssigment {
	_ = "STUB: not implemented"
	return *new(ColumnAssigment)
}

// RangeColumn creates named range column.
func RangeColumn[T Expression](name string) ColumnRange[T] { _ = "STUB: not implemented"; return nil }
