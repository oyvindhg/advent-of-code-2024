package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	area := parseInput("input/test.txt")
	result := solveFirst(area)

	want := 41

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	area := parseInput("input/test.txt")
	result := solveSecond(area)

	want := 6

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
