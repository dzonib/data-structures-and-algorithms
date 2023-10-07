package main

import "testing"

func TestIntegerReversal(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		got, _ := reverseAnInt(500)

		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		got, _ := reverseAnInt(-532)

		want := -235

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("positive 2", func(t *testing.T) {
		got, _ := reverseAnInt(512)

		want := 215

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("zero", func(t *testing.T) {
		got, _ := reverseAnInt(0)

		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
