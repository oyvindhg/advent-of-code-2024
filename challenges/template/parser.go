package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func parseInput(inputFile string) []string {
	_, currentPath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Error getting current path")
	}

	currentDir := filepath.Dir(currentPath)
	inputFilePath := filepath.Join(currentDir, inputFile)

	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("Error during file parsing: %v", err)
	}
	return strings.Split(string(data), "\n")
}
