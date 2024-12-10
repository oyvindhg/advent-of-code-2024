package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	heightMap := parseInput("input/test.txt")
	result := solveFirst(heightMap)

	want := 36

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	heightMap := parseInput("input/test.txt")
	result := solveSecond(heightMap)

	want := 0

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
