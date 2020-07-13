package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

// Hello function exposed for testing purposes
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {

	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

// main function logic
func main() {
	fmt.Println(Hello("Andreas and Dariusz", ""))
}
