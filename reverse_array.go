package main

import "fmt"

// Function to reverse an array
func reverseArray(arr []int) []int {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        // Swap elements at positions i and n-i-1
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
    return arr
}

func main() {
    var n int

    // Ask the user for the size of the array
    fmt.Print("Enter the size of the array: ")
    fmt.Scan(&n)

    // Create an array of size n
    arr := make([]int, n)

    // Get array elements from the user
    fmt.Println("Enter the elements of the array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    // Reverse the array
    reversedArr := reverseArray(arr)

    // Print the reversed array
    fmt.Println("Reversed array:", reversedArr)
}
