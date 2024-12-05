package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	printerInfo := parseInput("input/test.txt")
	result := solveFirst(printerInfo)

	want := 143

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	data := parseInput("input/test.txt")
	result := solveSecond(data)

	want := 123

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
