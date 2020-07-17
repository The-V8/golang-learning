package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	b.Run("repeat 5 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 5)
		}
	})

	b.Run("repeat 5000 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 5000)
		}
	})
}

func BenchmarkLibraryRepeat(b *testing.B) {
	b.Run("repeat 5 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Repeat("a", 5)
		}
	})

	b.Run("repeat 5000 times", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Repeat("a", 5000)
		}
	})
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
