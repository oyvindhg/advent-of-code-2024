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

func parseInput(inputFile string) []CalibrationEquation {
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

	calibrationEquations := []CalibrationEquation{}
	for scanner.Scan() {
		idWithEquationText := strings.Split(scanner.Text(), ": ")
		id, _ := strconv.Atoi(idWithEquationText[0])
		numberTextList := strings.Split(idWithEquationText[1], " ")
		numbers := []int{}
		for _, numberText := range numberTextList {
			number, _ := strconv.Atoi(numberText)
			numbers = append(numbers, number)
		}
		calibrationEquations = append(calibrationEquations, CalibrationEquation{result: id, numbers: numbers})
	}

	return calibrationEquations
}
