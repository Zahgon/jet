package mysql

import (
	"regexp"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
)

type unitType string

// List of interval unit types for MySQL
const (
	MICROSECOND        unitType = "MICROSECOND"
	SECOND             unitType = "SECOND"
	MINUTE             unitType = "MINUTE"
	HOUR               unitType = "HOUR"
	DAY                unitType = "DAY"
	WEEK               unitType = "WEEK"
	MONTH              unitType = "MONTH"
	QUARTER            unitType = "QUARTER"
	YEAR               unitType = "YEAR"
	SECOND_MICROSECOND unitType = "SECOND_MICROSECOND"
	MINUTE_MICROSECOND unitType = "MINUTE_MICROSECOND"
	MINUTE_SECOND      unitType = "MINUTE_SECOND"
	HOUR_MICROSECOND   unitType = "HOUR_MICROSECOND"
	HOUR_SECOND        unitType = "HOUR_SECOND"
	HOUR_MINUTE        unitType = "HOUR_MINUTE"
	DAY_MICROSECOND    unitType = "DAY_MICROSECOND"
	DAY_SECOND         unitType = "DAY_SECOND"
	DAY_MINUTE         unitType = "DAY_MINUTE"
	DAY_HOUR           unitType = "DAY_HOUR"
	YEAR_MONTH         unitType = "YEAR_MONTH"
)

// Interval is representation of MySQL interval
type Interval = jet.Interval

// INTERVAL creates new temporal interval.
//
//	In a case of MICROSECOND, SECOND, MINUTE, HOUR, DAY, WEEK, MONTH, QUARTER, YEAR unit type
//	value parameter has to be a number.
//			INTERVAL(1, DAY)
//	In a case of other unit types, value should be string with appropriate format.
//			INTERVAL("10:08:50", HOUR_SECOND)
func INTERVAL(value interface{}, unitType unitType) Interval {
	_ = "STUB: not implemented"
	return *new(Interval)
}

// INTERVALe creates new temporal interval from expresion and unit type.
func INTERVALe(expr Expression, unitType unitType) Interval {
	_ = "STUB: not implemented"
	return *new(Interval)
}

// INTERVALd creates new temporal interval from time.Duration
func INTERVALd(duration time.Duration) Interval { _ = "STUB: not implemented"; return *new(Interval) }

var (
	regexSecondMicrosecond = regexp.MustCompile(`^-?\d{1,2}\.\d+$`)                //'SECONDS.MICROSECONDS'
	regexMinuteMicrosecond = regexp.MustCompile(`^-?\d{1,2}:\d{2}\.\d+$`)          //'MINUTE:SECONDS.MICROSECONDS'
	regexMinuteSecond      = regexp.MustCompile(`^-?\d{1,2}:\d{2}$`)               //'MINUTE:SECONDS'
	regexHourMicrosecond   = regexp.MustCompile(`^-?\d{1,2}:\d{2}:\d{2}\.\d+$`)    //'HOUR:MINUTE:SECONDS.MICROSECONDS'
	regexHourSecond        = regexp.MustCompile(`^-?\d{1,2}:\d{2}:\d{2}$`)         //'HOUR:MINUTE:SECONDS'
	regexHourMinute        = regexp.MustCompile(`^-?\d{1,2}:\d{2}$`)               //'HOUR:MINUTE'
	regexDayMicrosecond    = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}:\d{2}.\d+$`) //'DAY HOUR:MINUTE:SECONDS'
	regexDaySecond         = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}:\d{2}$`)     //'DAY HOUR:MINUTE:SECONDS'
	regexDayMinute         = regexp.MustCompile(`^-?\d+ \d{1,2}:\d{2}$`)           //'DAY HOUR:MINUTE'
	regexDayHour           = regexp.MustCompile(`^-?\d+ \d{1,2}$`)                 //'DAY HOUR:MINUTE'
	regexYearMonth         = regexp.MustCompile(`^-?\d+-\d{1,2}$`)                 //'YEAR-MONTH'
)

func isNumericType(value interface{}) bool { _ = "STUB: not implemented"; return false }
