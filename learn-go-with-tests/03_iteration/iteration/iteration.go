package iteration

// Repeat repeats a string 5 times
func Repeat(s string, n int) string {
	output := ""
	for i := 0; i < n; i++ {
		output += s
	}
	return output
}
