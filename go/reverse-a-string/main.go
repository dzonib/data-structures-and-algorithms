package main

func main() {
	reverse("tesottar")
}

func reverse(s string) string {

	reversed := make([]byte, len(s))

	for i, _ := range s {
		// reversedIndex := i + len(s) - (i * 2) - 1   ??????
		reversedIndex := len(s) - 1 - i
		reversed[i] = s[reversedIndex]
	}
	return string(reversed)
}

//func reverse(s string) string {
//	rns := []rune(s) // convert to rune
//	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
//
//		// swap the letters of the string,
//		// like first with last and so on.
//		rns[i], rns[j] = rns[j], rns[i]
//	}
//
//	// return the reversed string.
//	return string(rns)
//}
//
