package postgres

import (
	"time"
)

type quantityAndUnit = float64
type unit = float64

// Interval unit types
const (
	YEAR unit = 123456789 + iota
	MONTH
	WEEK
	DAY
	HOUR
	MINUTE
	SECOND
	MILLISECOND
	MICROSECOND
	DECADE
	CENTURY
	MILLENNIUM
)

// INTERVAL creates new interval expression from the list of quantity-unit pairs.
//
//	INTERVAL(1, DAY, 3, MINUTE)
func INTERVAL(quantityAndUnit ...quantityAndUnit) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

// #nosec G602 -- false positive guarded by even-length check above and i+1 < n in loop

// INTERVALd creates interval expression from time.Duration
func INTERVALd(duration time.Duration) IntervalExpression {
	_ = "STUB: not implemented"
	return *new(IntervalExpression)
}

func unitToString(unit quantityAndUnit) string { _ = "STUB: not implemented"; return "" }

// additional field units for EXTRACT function

//---------------------------------------------------//
