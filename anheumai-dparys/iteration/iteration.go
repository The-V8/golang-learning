package iteration

const repCount = 5

// Repeater function
func Repeat(input string) string {
	var repeated string
	for i := 0; i < repCount; i++ {
		repeated += input
	}
	return repeated
}
