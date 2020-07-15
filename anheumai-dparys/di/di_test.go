package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Dariusz")

	got := buffer.String()
	want := "Hello, Dariusz"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
