package main

import (
	"math"
	"slices"
)

func solveFirst(input Locations) int {
	left := append([]int{}, input.left...)
	right := append([]int{}, input.right...)

	slices.Sort(left)
	slices.Sort(right)

	sumDifference := 0

	for i, leftLocation := range left {
		rightLocation := right[i]
		sumDifference += int(math.Abs(float64(leftLocation - rightLocation)))
	}

	return sumDifference
}

func solveSecond(input Locations) int {
	rightCounter := make(map[int]int)
	for _, rightLocation := range input.right {
		rightCounter[rightLocation] += 1
	}

	similarityScore := 0
	for _, leftLocation := range input.left {
		similarityScore += leftLocation * rightCounter[leftLocation]
	}

	return similarityScore
}
