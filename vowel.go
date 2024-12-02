package main

import (
	"fmt"
	"strings"
)

// Function to check if a character is a vowel
func isVowel(ch rune) bool {
	// Convert character to lowercase for case-insensitivity
	ch = rune(strings.ToLower(string(ch))[0])

	// Check if the character is in the set of vowels
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func main() {
	var char string

	// Ask the user to input a character
	fmt.Print("Enter a character: ")
	fmt.Scan(&char)

	// Validate input to ensure it is a single character
	if len(char) != 1 {
		fmt.Println("Please enter a single character.")
		return
	}

	// Check if the character is a vowel
	if isVowel(rune(char[0])) {
		fmt.Printf("'%s' is a vowel.\n", char)
	} else {
		fmt.Printf("'%s' is not a vowel.\n", char)
	}
}
