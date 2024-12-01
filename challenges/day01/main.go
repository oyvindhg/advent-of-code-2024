package main

import "fmt"

func main() {
	data := parseInput("input/input.txt")
	fmt.Println("Solution 1: ", solveFirst(data))
	fmt.Println("Solution 2: ", solveSecond(data))
}
