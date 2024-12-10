package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func parseInput(inputFile string) [][]int {
	_, currentPath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Error getting current path")
	}

	currentDir := filepath.Dir(currentPath)
	inputFilePath := filepath.Join(currentDir, inputFile)

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Error when reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heightMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		heightRow := []int{}
		for _, heightRune := range line {
			heightRow = append(heightRow, int(heightRune-'0'))
		}
		heightMap = append(heightMap, heightRow)
	}

	return heightMap
}
