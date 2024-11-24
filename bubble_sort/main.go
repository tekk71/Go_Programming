package main

import (
    "fmt"
)

// BubbleSort sorts a slice of integers using the Bubble Sort algorithm.
func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // Swap
            }
        }
    }
}

func main() {
    // Example array to sort
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Unsorted array:", arr)

    // Sort the array
    BubbleSort(arr)
    fmt.Println("Sorted array:  ", arr)
}
