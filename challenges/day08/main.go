package main

import (
	"fmt"
	"time"
)

func main() {
	cityMap := parseInput("input/input.txt")

	start := time.Now()
	fmt.Println("--Task 1--")
	fmt.Println("Solution:", solveFirst(cityMap))
	fmt.Printf("Time: %v ms\n\n", time.Since(start).Milliseconds())

	start = time.Now()
	fmt.Println("--Task 2--")
	fmt.Println("Solution:", solveSecond(cityMap))
	fmt.Printf("Time: %v ms\n", time.Since(start).Milliseconds())
}
