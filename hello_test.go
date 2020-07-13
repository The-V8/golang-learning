package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Andreas", "")
		expected := "Hello, Andreas"
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Say hello to Dariusz", func(t *testing.T) {
		got := Hello("Dariusz", "")
		expected := "Hello, Dariusz"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Claude", "French")
		want := "Bonjour, Claude"
		assertCorrectMessage(t, got, want)
	})

}
