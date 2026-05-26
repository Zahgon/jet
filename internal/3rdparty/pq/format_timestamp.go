package pq

// Copyright (c) 2011-2013, 'pq' Contributors Portions Copyright (C) 2011 Blake Mizerany

import (
	"time"
)

// FormatTimestamp formats t into Postgres' text format for timestamps. From: github.com/lib/pq
func FormatTimestamp(t time.Time) []byte {
	_ = "STUB: not implemented"
	// Need to send dates before 0001 A.D. with " BC" suffix, instead of the
	// minus sign preferred by Go.
	// Beware, "0000" in ISO is "1 BC", "-0001" is "2 BC" and so on
	return nil
}

// flip year sign, and add 1, e.g: "0" will be "1", and "-10" will be "11"

// RFC3339Nano already printed the minus sign
