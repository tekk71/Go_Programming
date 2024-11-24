package main

import (
    "fmt"
    "strings"
)

// BubbleSort sorts a slice of strings alphabetically using the Bubble Sort algorithm.
func BubbleSort(files []string) {
    n := len(files)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if strings.ToLower(files[j]) > strings.ToLower(files[j+1]) {  // Compare alphabetically
                files[j], files[j+1] = files[j+1], files[j]  // Swap
            }
        }
    }
}

func main() {
    // Example array of file names
    files := []string{
"Tiger", "House", "Ice", "Frog", "Xylophone", "Zebra", "Lemon", "Ocean", "Penguin", "Giraffe", "Rabbit", "Apple", "Queen", "Cat", "Dog", "Umbrella", "Mountain", "Snake", "Violin", "Kite", "Banana", "Jaguar", "Yacht", "Elephant", "Nest", "Whale",
    }

    fmt.Println("Unsorted file names:", files)

    // Sort the file names
    BubbleSort(files)
    fmt.Println("Sorted file names:  ", files)
}
