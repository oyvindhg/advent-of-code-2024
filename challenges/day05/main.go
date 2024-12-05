package main

import (
	"fmt"
	"time"
)

func main() {
	printerInfo := parseInput("input/input.txt")

	start := time.Now()
	fmt.Println("--Task 1--")
	fmt.Println("Solution:", solveFirst(printerInfo))
	fmt.Printf("Time: %v ms\n\n", time.Since(start).Milliseconds())

	start = time.Now()
	fmt.Println("--Task 2--")
	fmt.Println("Solution:", solveSecond(printerInfo))
	fmt.Printf("Time: %v ms\n", time.Since(start).Milliseconds())
}
