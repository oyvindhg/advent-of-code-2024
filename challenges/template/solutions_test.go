package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	data := parseInput("input/input.txt")

	var want int = 0
	result := solveFirst(data)
	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	data := parseInput("input/input.txt")

	var want int = 0
	result := solveFirst(data)
	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
