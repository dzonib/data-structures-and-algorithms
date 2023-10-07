package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Coding Exercise - Integer Reversal
//
//Given an integer, return an integer that is the reverse ordering of numbers.
//
//	Examples:
//
//		reverseInt(15) === 51;
//		reverseInt(981) === 189;
//		reverseInt(500) === 5;
//		reverseInt(-15) === -51;
//		reverseInt(-90) === -9;

func main() {
	reversedInt, err := reverseAnInt(-40012)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("RESULT", reversedInt)
}

func reverseAnInt(i int) (int, error) {
	isPositive := math.Signbit(float64(i))

	// int to str
	numStr := strconv.Itoa(i)

	// Split the string into an array of letters
	letters := strings.Split(numStr, "")

	if isPositive {
		// take all except first element (-)
		letters = letters[1:]
	}

	// No, the code you've provided will not remove the first element from the letters slice. Instead, it will create a
	// new slice that includes all elements except the first element of the original letters slice.
	// The line letters = append(letters[1:]) is using slicing to create a new slice starting from the second element
	// of the original letters slice to the end, and then assigning this new slice back to the letters variable.
	// This effectively creates a new slice that excludes the first element of the original slice, but it doesn't
	// remove the element from the original slice itself.
	//letters = append(letters[1:])

	// swap x to the end ?
	//x, y := letters[0], letters[1:]
	//a = append(a, x)

	// 40012 len == 5

	for i, n := 0, len(letters)-1; i < n; i, n = i+1, n-1 {
		// loop 1
		// i == 0, n == 4
		// i < n == true

		// loop 2
		// i == 1, n == 3
		// i < n == true

		//loop 3
		//i == 2, n == 2
		// i < n == false loop does not run

		letters[i], letters[n] = letters[n], letters[i]
		// loop 1
		// 20014
		// loop 2
		// 21004
	}

	// https://www.geeksforgeeks.org/math-signbit-function-in-golang-with-examples/
	// returns true or false depending on if it is a positive or negative number

	// Concatenate the letters into a single string
	letterStr := strings.Join(letters, "")

	num, err := strconv.Atoi(letterStr)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	if isPositive {
		return num * -1, nil
	} else {
		return num * 1, nil
	}

}
