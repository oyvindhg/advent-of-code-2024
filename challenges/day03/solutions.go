package main

import (
	"regexp"
	"strconv"
	"strings"
)

func solveFirst(program string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	multiplications := re.FindAllString(program, -1)

	multiplicationSum := 0
	for _, multiplication := range multiplications {
		numbers := strings.Split(multiplication[4:len(multiplication)-1], ",")
		leftNumber, _ := strconv.Atoi(numbers[0])
		rightNumber, _ := strconv.Atoi(numbers[1])

		multiplicationSum += leftNumber * rightNumber
	}

	return multiplicationSum
}

func solveSecond(program string) int {
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don't\(\)|do\(\)`)

	instructions := re.FindAllString(program, -1)

	multiplicationSum := 0
	isEnabled := true
	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			isEnabled = true
		case "don't()":
			isEnabled = false
		default:
			if isEnabled {
				numbers := strings.Split(instruction[4:len(instruction)-1], ",")
				leftNumber, _ := strconv.Atoi(numbers[0])
				rightNumber, _ := strconv.Atoi(numbers[1])

				multiplicationSum += leftNumber * rightNumber
			}
		}
	}

	return multiplicationSum
}
