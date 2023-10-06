package integerreversal

import "testing"

func TestIntegerReversal(t *testing.T) {
	got := reverseAnInt(500)

	want := 005

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

