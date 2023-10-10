package anagram

import "testing"

func TestAnagram(t *testing.T) {
	t.Run("anagram one", func(t *testing.T) {
		got := IsAnagram("listen", "silent")

		want := true

		assertCorrectMessage(t, got, want)

	})

	t.Run("anagram two", func(t *testing.T) {
		got := IsAnagram("listen", "silent")

		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("not an anagram", func(t *testing.T) {
		got := IsAnagram("ooooo", "silent")

		want := false

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	//	t.Helper() is needed to tell the test suite that this method is a helper.
	//  By doing this when it fails, the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
