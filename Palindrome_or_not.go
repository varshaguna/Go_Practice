package main

import (
	"fmt"
)

func main() {
	var num, reversed, remainder, original int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	original = num // Store the original number to compare later
	reversed = 0   // Initialize reversed number as 0

	// Reverse the number
	for num != 0 {
		remainder = num % 10        // Extract the last digit
		reversed = reversed*10 + remainder // Build the reversed number
		num = num / 10              // Remove the last digit
	}

	// Check if the original number is equal to its reverse
	if original == reversed {
		fmt.Printf("The number %d is a palindrome.\n", original)
	} else {
		fmt.Printf("The number %d is not a palindrome.\n", original)
	}
}
