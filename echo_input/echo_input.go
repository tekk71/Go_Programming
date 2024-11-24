package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Create a reader to read input from standard input (keyboard)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter some text: ")

    // Read the input until the newline character
    input, _ := reader.ReadString('\n')

    // Print the input back to the user
    fmt.Println("You entered:", input)
}
