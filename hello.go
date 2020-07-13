package main

import "fmt"

const (
	helloPrefix = "Hello, "
)

// Hello function exposed for testing purposes
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name + "."
}

// main function logic
func main() {
	fmt.Println(Hello("Andreas and Dariusz"))
}
