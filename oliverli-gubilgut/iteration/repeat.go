package iteration

func Repeat(character string, ntimes int) string {
	var repeated string
	for i := 0; i < ntimes; i++ {
		repeated += character
	}
	return repeated
}
