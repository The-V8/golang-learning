package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Oliver")

	got := buffer.String()
	want := "Hello, Oliver"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
