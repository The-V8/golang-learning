package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const helloPostfix = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + helloPostfix
}

func main() {
	fmt.Println(Hello("world", ""))
}
