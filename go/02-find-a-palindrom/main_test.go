package main

import "testing"

func TestPalindrom(t *testing.T) {
	//	t.Run()

	got := palindrome("Abba")

	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}

}
