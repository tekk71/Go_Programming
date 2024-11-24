package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open the input file
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()

    // Create the output file
    outputFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    // Create a scanner to read the input file
    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(outputFile)

    // Process each line and write to the output file
    for scanner.Scan() {
        line := scanner.Text()
        processedLine := strings.ToUpper(line) // Example processing: convert to uppercase
        fmt.Fprintln(writer, processedLine)
    }

    // Check for any scanner errors
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }

    // Flush the writer to ensure all data is written to the output file
    writer.Flush()

    fmt.Println("File processing completed. Check output.txt for results.")
}
