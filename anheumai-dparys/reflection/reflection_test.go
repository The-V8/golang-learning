package reflection

import (
	"reflect"
	"testing"
)

var numberOfCalls = 1

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Dariusz"},
			[]string{"Dariusz"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Dariusz", "Stuttgart"},
			[]string{"Dariusz", "Stuttgart"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Dariusz", 99},
			[]string{"Dariusz"},
		},
		{
			"Nested anonymous Structs",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Dariusz", struct {
				Age  int
				City string
			}{99, "Stuttgart"}},
			[]string{"Dariusz", "Stuttgart"},
		},
		{
			"Nested Structs",
			Person{
				"Dariusz",
				Profile{99, "Stuttgart"},
			},
			[]string{"Dariusz", "Stuttgart"},
		},
		{
			"Using Pointers",
			&Person{
				"Dariusz",
				Profile{99, "Stuttgart"},
			},
			[]string{"Dariusz", "Stuttgart"},
		},
		{
			"Slices",
			[]Profile{
				{99, "Stuttgart"},
				{100, "Seattle"},
			},
			[]string{"Stuttgart", "Seattle"},
		},
		{
			"Arrays",
			[1]Profile{
				{99, "Stuttgart"},
			},
			[]string{"Stuttgart"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]Profile{
			"Dariusz": {99, "Stuttgart"},
			"Satya":   {100, "Seattle"},
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Stuttgart")
		assertContains(t, got, "Seattle")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{99, "Stuttgart"}
			ch <- Profile{100, "Seattle"}
			close(ch)
		}()

		var got []string
		want := []string{"Stuttgart", "Seattle"}

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		fn := func() (Profile, Profile) {
			return Profile{99, "Stuttgart"}, Profile{100, "Seattle"}
		}

		var got []string
		want := []string{"Stuttgart", "Seattle"}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, list []string, term string) {
	t.Helper()

	contains := false
	for _, x := range list {
		if x == term {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but no", list, term)
	}
}
