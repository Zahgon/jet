package datetime

import (
	//"github.com/go-jet/jet/v2/internal/utils/min"
	"time"
)

// ExtractTimeComponents extracts number of days, hours, minutes, seconds, microseconds from duration
func ExtractTimeComponents(duration time.Duration) (days, hours, minutes, seconds, microseconds int64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0
}

// TryParseAsTime attempts to parse the provided value as a time using one of the given formats.
//
// The function iterates over the provided formats and tries to parse the value into a time.Time object.
// It returns the parsed time and a boolean indicating whether the parsing was successful.
func TryParseAsTime(value interface{}, formats []string) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// sqlite
