package internal

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

var (
	errCastOverFlow = fmt.Errorf("cannot cast a negative value to an unsigned value, buffer overflow error")
)

// NullBool struct
type NullBool struct {
	sql.NullBool
}

// Scan implements the Scanner interface.
func (nb *NullBool) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// NullTime struct
type NullTime struct {
	sql.NullTime
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Some of the drivers (pgx, mysql) are not parsing all of the time formats(date, time with time zone,...) and are just forwarding string value.
// At this point we try to parse those values using some of the predefined formats

// sqlite
// go-sql-driver/mysql
// pgx
// pgx

// NullUInt64 struct
type NullUInt64 struct {
	UInt64 uint64
	Valid  bool
}

// Scan implements the Scanner interface.
func (n *NullUInt64) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Value implements the driver Valuer interface.
func (n NullUInt64) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
