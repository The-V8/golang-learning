package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("The-V8")
	want := "Hello, The-V8!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
