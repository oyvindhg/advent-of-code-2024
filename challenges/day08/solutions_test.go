package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	cityMap := parseInput("input/test.txt")
	result := solveFirst(cityMap)

	want := 14

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	cityMap := parseInput("input/test.txt")
	result := solveSecond(cityMap)

	want := 34

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
