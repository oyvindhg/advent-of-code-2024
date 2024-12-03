package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	program := parseInput("input/test.txt")
	result := solveFirst(program)

	want := 161

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	program := parseInput("input/test.txt")
	result := solveSecond(program)

	want := 48

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
