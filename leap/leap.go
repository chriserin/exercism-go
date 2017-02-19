//Package leap provides functions to determine if a year is or is not in fact a leap year
package leap

const testVersion = 3

// Returns true if input year is a leap year.
func IsLeapYear(x int) bool {
	return x%4 == 0 && (x%100 != 0 || x%400 == 0)
}
