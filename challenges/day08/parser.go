package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func parseInput(inputFile string) [][]rune {
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

	cityMap := [][]rune{}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		cityMap = append(cityMap, line)
	}

	return cityMap
}
