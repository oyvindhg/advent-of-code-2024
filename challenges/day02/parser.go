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

	scanner := bufio.NewScanner(file)

	reports := [][]int{}
	for scanner.Scan() {
		split_line := strings.Split(scanner.Text(), " ")
		report := []int{}
		for _, levelText := range split_line {
			level, _ := strconv.Atoi(levelText)
			report = append(report, level)
		}
		reports = append(reports, report)
	}

	return reports
}
