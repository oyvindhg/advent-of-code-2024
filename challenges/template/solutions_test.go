package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	data := parseInput("input/test.txt")
	result := solveFirst(data)

	want := 0

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	data := parseInput("input/test.txt")
	result := solveSecond(data)

	want := 0

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
