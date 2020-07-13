package main

import "fmt"

// Hello function exposed for testing purposes
func Hello() string {
	return "Hello, world."
}

// main finction logic
func main() {
	fmt.Println(Hello())
}
