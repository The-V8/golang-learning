package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloPostfix = "!"

func Hello(name string) string {
	return englishHelloPrefix + name + englishHelloPostfix
}

func main() {
	fmt.Println(Hello("world"))
}
