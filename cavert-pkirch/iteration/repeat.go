package iteration

// Returns a string with the given string repeated by the given number.
func Repeat(char string, iterations int) string {
	var repeated string

	for i := 0; i < iterations; i++ {
		repeated += char
	}

	return repeated
}
