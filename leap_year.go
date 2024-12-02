package main

import "fmt"

// Function to check if a year is a leap year
func isLeapYear(year int) bool {
    // A year is a leap year if:
    // 1. It is divisible by 4, and
    // 2. It is NOT divisible by 100, unless it is also divisible by 400
    if year%4 == 0 {
        if year%100 == 0 {
            // If divisible by 100, it must also be divisible by 400 to be a leap year
            return year%400 == 0
        }
        return true // Divisible by 4 but not by 100
    }
    return false // Not divisible by 4
}

func main() {
    var year int

    // Get the year input from the user
    fmt.Print("Enter a year: ")
    fmt.Scan(&year)

    // Check if the year is a leap year and print the result
    if isLeapYear(year) {
        fmt.Printf("%d is a leap year.\n", year)
    } else {
        fmt.Printf("%d is not a leap year.\n", year)
    }
}
