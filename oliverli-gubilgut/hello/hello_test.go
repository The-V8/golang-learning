package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	//testing comment --- subtest 1
	t.Run("saying Hi to Space Nation", func(t *testing.T) {
		got := Hello("dear all", "")
		want := "Hello, dear all"

		assertCorrectMessage(t, got, want)
	})

	//subtest 2
	t.Run("Hello with blank completed as to default", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Asgardia"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Espanol", func(t *testing.T) {
		got := Hello("Primo", "Spanish")
		want := "Hola, Primo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Francais", func(t *testing.T) {
		got := Hello("Camarade", "French")
		want := "Bonjour, Camarade"
		assertCorrectMessage(t, got, want)
	})
}
