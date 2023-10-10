package mostchar

import "testing"

func TestMostChar(t *testing.T) {
	t.Run("first letters repeat", func(t *testing.T) {
		got := MostChar("aaatest")

		want := "a"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("mixed letters", func(t *testing.T) {
		got := MostChar("btebstawrybas13  421 bbasbb23-")

		want := "b"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
