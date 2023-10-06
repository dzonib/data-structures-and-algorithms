package main

import (
	"strings"
)

// Coding Exercise - Palindromes
//
//	Given a string, return true if the string is a palindrome or false if it is not.
//	Palindromes are strings that form the same word if it is reversed.
//
//	*Do* include spaces and punctuation in determining if the string is a palindrome.
//
//	Examples:
//
//		palindrome("abba") === true
//		palindrome("abcdefg") === false
func main() {
	palindrome("Test")
}
func palindrome(s string) bool {
	rns := []rune(s) // convert to rune

	// oop that uses two variables, i and j, to traverse the characters of the string from both ends toward the center.
	// i starts at the beginning (0), and j starts at the end (the last index of the rns slice).
	// The loop continues as long as i is less than j.

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// 0, 3
		// swap T t, t T

		// 1, 2
		// swap e s, s e

		// rns[i], rns[j] = rns[j], rns[i]: This line swaps the characters at positions i and j in the rns slice.
		// It effectively reverses the order of characters in the string.
		rns[i], rns[j] = rns[j], rns[i]
	}

	if strings.ToLower(string(rns)) == strings.ToLower(s) {
		return true
	} else {
		return false
	}
}
