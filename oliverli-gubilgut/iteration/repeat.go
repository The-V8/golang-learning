package iteration

import "strings"

// Repeat implements a custom string Repeat function
func Repeat(character string, ntimes int) string {
	var repeated string
	for i := 0; i < ntimes; i++ {
		repeated += character
	}
	return repeated
}

// RepeatStringLib uses the standard lib strings Repeat funct
func RepeatStringLib(character string, ntimes int) string {
	return strings.Repeat(character, ntimes)
}
