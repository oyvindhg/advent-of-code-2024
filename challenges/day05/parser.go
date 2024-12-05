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

func parseInput(inputFile string) PrinterInfo {
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

	pageOrderingRules := []PageOrderingRule{}
	updates := []Update{}
	isParsingOrderingRules := true
	for scanner.Scan() {
		if isParsingOrderingRules {
			if scanner.Text() == "" {
				isParsingOrderingRules = false
			} else {
				rule := strings.Split(scanner.Text(), "|")
				from, _ := strconv.Atoi(rule[0])
				to, _ := strconv.Atoi(rule[1])
				pageOrderingRules = append(pageOrderingRules, PageOrderingRule{from: from, to: to})
			}
		} else {
			pageLine := strings.Split(scanner.Text(), ",")
			pages := []int{}
			for _, pageText := range pageLine {
				page, _ := strconv.Atoi(pageText)
				pages = append(pages, page)
			}
			updates = append(updates, Update{pages: pages})
		}
	}

	return PrinterInfo{pageOrderingRules: pageOrderingRules, updates: updates}
}
