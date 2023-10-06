package main

import "testing"

//Coding Exercise - String Reversal
//
//	Given a string, return a new string with the reversed order of characters.
//
//	Examples:
//
//		reverse('apple') === 'leppa'
//		reverse('hello') === 'olleh'
//		reverse('Greetings!') === '!sgniteerG'
//
//		You should put your implementation into the reverse function that was defined for you below.  You don't need to call reverse yourself.
//		Example of what your completed code should look like:
//
//		function reverse (str) {
//			// Code to implement the reverse function
//		}

func TestReverse(t *testing.T) {
	got := reverse("Hello")

	want := "olleH"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
