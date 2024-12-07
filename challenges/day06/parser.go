package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func parseInput(inputFile string) [][]rune {
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

	text_split := strings.Split(string(data), "\n")

	runes := [][]rune{}
	for _, text_line := range text_split {
		rune_line := []rune(text_line)
		runes = append(runes, rune_line)
	}

	return runes
}
