package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collections of unspecified length", func(t *testing.T) {
		numbers := []int{5, 66, 9}

		got := Sum(numbers)
		want := 80

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
