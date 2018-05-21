package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%x'.", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(4, 6)
	fmt.Println(sum)
	// Output: 10
}
