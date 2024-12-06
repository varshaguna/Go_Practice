package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binaryStr string
	fmt.Print("Enter the binary number: ")
	fmt.Scanln(&binaryStr)

	decimal, err := strconv.ParseInt(binaryStr, 2, 64)

	if err != nil {
		fmt.Println("Invalid binary number", err)
		return
	}
	fmt.Printf("The decimal equivalent of binary %s is %d\n", binaryStr, decimal)
}
