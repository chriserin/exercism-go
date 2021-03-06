// Package clause.
package gigasecond

import (
	"time"
)

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	duration, _ := time.ParseDuration("1000000000s")
	return t.Add(duration)
}
