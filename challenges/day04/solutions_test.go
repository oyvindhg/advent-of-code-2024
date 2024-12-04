package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	puzzle := parseInput("input/test.txt")
	result := solveFirst(puzzle)

	want := 18

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	puzzle := parseInput("input/test.txt")
	result := solveSecond(puzzle)

	want := 9

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
