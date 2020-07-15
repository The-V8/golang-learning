package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "number 1"}

	t.Run("Known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "number 1"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("tbd")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		term := "test"
		definition := "this is just a test"
		err := dictionary.Add(term, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}
		err := dictionary.Add(term, definition)

		assertError(t, err, ErrTermExists)
		assertDefinition(t, dictionary, term, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("term exists", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}
		newDefinition := "some new stuff"

		err := dictionary.Update(term, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, newDefinition)
	})

	t.Run("term does not exists", func(t *testing.T) {
		term := "tbd"
		dictionary := Dictionary{}
		newDefinition := "some new stuff"

		err := dictionary.Update(term, newDefinition)

		assertError(t, err, ErrTermDoesNotExists)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Delete existing", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}

		dictionary.Delete(term)

		_, err := dictionary.Search(term)
		if err != ErrNotFound {
			t.Errorf("Expected %v to be deleted", term)
		}
	})
}

func assertDefinition(t *testing.T, d Dictionary, term, definition string) {
	t.Helper()

	got, err := d.Search(term)
	if err != nil {
		t.Fatal("where is the word, bro?", err)
	}

	if definition != got {
		t.Errorf("want %v, got %v", definition, got)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
