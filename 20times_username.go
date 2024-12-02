package main

import "fmt"

func main() {
	var username string

	// Ask the user for their username
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)

	// Print the username 20 times
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d: %s\n", i, username)
	}
}
