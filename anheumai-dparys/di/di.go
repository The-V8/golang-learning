package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Dariusz")
}

func main() {
	http.ListenAndServe(":4243", http.HandlerFunc(MyGreetHandler))
}
