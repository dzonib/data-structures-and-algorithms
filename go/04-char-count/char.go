package mostchar

//Coding Exercise - Max Chars
//
//Given a string, return the character that is most commonly used in the string.
//
//Examples:
//
//maxChar("abcccccccd") === "c"
//maxChar("apple 1231111") === "1"

func MostChar(s string) string {
	// final result
	var result string

	// create a placeholder (string) for the highest value in the map placeholder
	maxValue := 0

	// split string into a rune array
	runeArr := []rune(s)

	// create a placeholder (map) for chars and value (string[int])
	m := make(map[string]int)

	// loop over a rune array and assign the chars to map
	for _, c := range runeArr {
		value, exists := m[string(c)]
		if !exists {
			m[string(c)] = 1
		} else {
			m[string(c)] = value + 1
		}
	}

	// read the values from placeholder and assign the highest value to the maxValue
	// loop over the map
	for key, value := range m {
		if value > maxValue {
			maxValue = value
			result = key
		}
	}

	return result
}
