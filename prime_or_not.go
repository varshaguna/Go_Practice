package main

import "fmt"

// Function to check if a number is prime
func isPrime(n int) bool {
	// Handle special cases
	if n <= 1 {
		return false // 1 and numbers less than 1 are not prime
	}

	// Loop through numbers from 2 to the square root of n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false // n is divisible by i, so it's not prime
		}
	}
	return true // If no divisors are found, it's prime
}

func main() {
	var num int

	// Ask the user to input a number
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is prime and print the result
	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}
