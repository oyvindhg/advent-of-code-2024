package main

import (
	"testing"
)

func TestSolveFirst(t *testing.T) {
	diskmap := parseInput("input/test.txt")
	result := solveFirst(diskmap)

	want := 1928

	if result != want {
		t.Fatalf("Solution 1 = %v, expected %v", result, want)
	}
}

func TestSolveSecond(t *testing.T) {
	diskmap := parseInput("input/test.txt")
	result := solveSecond(diskmap)

	want := 2858

	if result != want {
		t.Fatalf("Solution 2 = %v, expected %v", result, want)
	}
}
