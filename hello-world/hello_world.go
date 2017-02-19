// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package greeting should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package greeting

const testVersion = 3

func HelloWorld(a string) string {
	if a == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + a + "!"
	}
}
