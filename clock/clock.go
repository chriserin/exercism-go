// Package for formatting a clock struct
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

// Initialize clock
func New(hour, minute int) Clock {
	minutes := minute
	subtract_hours := 0

	for minutes < 0 {
		minutes += 60
		subtract_hours += 1
	}

	hours := hour + minutes/60
	hours -= subtract_hours

	for hours >= 24 {
		hours -= 24
	}

	for hours < 0 {
		hours += 24
	}
	return Clock{hours, minutes % 60}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(more_minutes int) Clock {
	minutes := clock.minute + more_minutes
	additional_hours := 0

	for minutes < 0 {
		minutes += 60
		additional_hours -= 1
	}

	for minutes > 60 {
		minutes -= 60
		additional_hours += 1
	}

	return Clock{hour: (24 + clock.hour + additional_hours) % 24, minute: minutes}
}
