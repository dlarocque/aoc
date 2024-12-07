package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	result := solve(string(data))

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %d ms\n", elapsed.Milliseconds()) // Output in milliseconds
}

func solve(input string) int {
	return 0
}
