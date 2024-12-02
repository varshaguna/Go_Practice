package main

import "fmt"

// Function to check if a number is prime
func isPrime(num int) bool {
    if num < 2 {
        return false // Numbers less than 2 are not prime
    }

    // Check if the number is divisible by any number from 2 to num-1
    for i := 2; i < num; i++ {
        if num%i == 0 {
            return false // Not a prime number if divisible
        }
    }
    return true // Prime number if not divisible
}

// Function to print all prime numbers from 1 to N
func printPrimes(n int) {
    fmt.Printf("Prime numbers between 1 and %d are:\n", n)

    for i := 1; i <= n; i++ {
        if isPrime(i) {
            fmt.Println(i) // Print the prime number
        }
    }
}

func main() {
    var n int

    // Get the value of N from the user
    fmt.Print("Enter a number N: ")
    fmt.Scan(&n)

    // Print all prime numbers between 1 and N
    printPrimes(n)
}
