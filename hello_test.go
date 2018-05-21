package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("ndi")
		want := "Hello, ndi"

		if got != want {
			t.Errorf("while testing Hello I expected %s, but got %s", want, got)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("while testing Hello I expected %s, but got %s", want, got)
		}
	})
}
