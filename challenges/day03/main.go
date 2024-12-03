package main

import (
	"fmt"
)

func main() {
	program := parseInput("input/input.txt")
	fmt.Println("Solution 1: ", solveFirst(program))
	fmt.Println("Solution 2: ", solveSecond(program))
}
