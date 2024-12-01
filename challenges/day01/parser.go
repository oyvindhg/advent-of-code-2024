package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func parseInput(inputFile string) Locations {
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

	locations := Locations{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split_line := strings.Split(scanner.Text(), "   ")
		first_location, _ := strconv.Atoi(split_line[0])
		second_location, _ := strconv.Atoi(split_line[1])

		locations.left = append(locations.left, first_location)
		locations.right = append(locations.right, second_location)
	}

	return locations
}
