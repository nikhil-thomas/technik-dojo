// Package documentme is for showing the documentaiton capabilities of go
// It is a native package
package documentme

// Pie is a global variable
// This is a silly comment
const Pie = 3.1415912

// S1 finds the length of a string
// it iterates over the string using range
func S1(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for range s {
		n++
	}
	return n
}

// F1 function returns the double value of its input integer
// A better name would have been Double()!
func F1(n int) int {
	return 2 * n
}
