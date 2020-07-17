package iteration

func Repeat(char string, iterations int) string {
	var repeated string

	for i := 0; i < iterations; i++ {
		repeated += char
	}

	return repeated
}
