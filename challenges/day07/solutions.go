package main

import (
	"strconv"
)

func isValidEquation(desiredResult int, numbers []int, currentId int, currentResult int, allowConcatenation bool) bool {
	if currentId == len(numbers) {
		return desiredResult == currentResult
	}
	if currentResult > desiredResult {
		return false
	}
	if currentId == 0 {
		return isValidEquation(desiredResult, numbers, currentId+1, numbers[0], allowConcatenation)
	} else {
		plusResult := currentResult + numbers[currentId]
		multiplyResult := currentResult * numbers[currentId]

		isPlusValid := isValidEquation(desiredResult, numbers, currentId+1, plusResult, allowConcatenation)
		isMultiplyValid := isValidEquation(desiredResult, numbers, currentId+1, multiplyResult, allowConcatenation)
		isConcatenationValid := false
		if allowConcatenation {
			concatenationResult, _ := strconv.Atoi(strconv.Itoa(currentResult) + strconv.Itoa(numbers[currentId]))
			isConcatenationValid = isValidEquation(desiredResult, numbers, currentId+1, concatenationResult, allowConcatenation)
		}
		return isPlusValid || isMultiplyValid || isConcatenationValid
	}
}

func solveFirst(calibrationEquations []CalibrationEquation) int {
	calibrationResult := 0

	for _, calibrationEquation := range calibrationEquations {
		if isValidEquation(calibrationEquation.result, calibrationEquation.numbers, 0, 0, false) {
			calibrationResult += calibrationEquation.result
		}
	}

	return calibrationResult
}

func solveSecond(calibrationEquations []CalibrationEquation) int {
	calibrationResult := 0

	for _, calibrationEquation := range calibrationEquations {
		if isValidEquation(calibrationEquation.result, calibrationEquation.numbers, 0, 0, true) {
			calibrationResult += calibrationEquation.result
		}
	}

	return calibrationResult
}
