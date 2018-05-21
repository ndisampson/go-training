package iteration

import (
	"fmt"
	"testing"
)

var repeatTests = []struct {
	in string
	n int
	out string
} {
	{"a", 5, "aaaaa"},
	{"b", 4, "bbbbb"}
}

func TestRepeat() {
	for _, tt := range repeatTests {
		t.Run(tt.in, func*(t *testing.T)) {
			r := Repeat(tt.in, tt.n)
			if r != tt.out {
				t.Errorf("got %q, want %q", r, tt.out)
			}
		}
	}
}

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeat != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	s := Repeat("a", 5)
	fmt.Println(s)
	// Output: aaaaa
}
