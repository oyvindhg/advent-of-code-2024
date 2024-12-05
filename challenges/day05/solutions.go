package main

func isValidUpdate(pageToPrecedingPages map[int][]int, update Update) bool {
	seenPrecedingPages := make(map[int]bool)
	for _, page := range update.pages {
		_, ok := seenPrecedingPages[page]
		if ok {
			return false
		}
		precedingPages, ok := pageToPrecedingPages[page]
		if ok {
			for _, precedingPage := range precedingPages {
				seenPrecedingPages[precedingPage] = true
			}
		}
	}
	return true
}

func solveFirst(printerInfo PrinterInfo) int {

	pageToPrecedingPages := make(map[int][]int)

	for _, pageOrderingRule := range printerInfo.pageOrderingRules {
		from := pageOrderingRule.from
		to := pageOrderingRule.to
		pageToPrecedingPages[to] = append(pageToPrecedingPages[to], from)
	}

	middlePagesSum := 0
	for _, update := range printerInfo.updates {
		isValid := isValidUpdate(pageToPrecedingPages, update)
		if isValid {
			middleIndex := len(update.pages) / 2
			middlePagesSum += update.pages[middleIndex]
		}
	}

	return middlePagesSum
}

func solveSecond(printerInfo PrinterInfo) int {
	pageToPrecedingPages := make(map[int][]int)
	pageToDependentPages := make(map[int][]int)

	for _, pageOrderingRule := range printerInfo.pageOrderingRules {
		from := pageOrderingRule.from
		to := pageOrderingRule.to
		pageToDependentPages[from] = append(pageToDependentPages[from], to)
		pageToPrecedingPages[to] = append(pageToPrecedingPages[to], from)
	}

	middlePagesSum := 0
	for _, update := range printerInfo.updates {
		isValid := isValidUpdate(pageToPrecedingPages, update)
		if !isValid {
			existingPages := make(map[int]bool)
			remainingPages := make(map[int]bool)
			for _, page := range update.pages {
				remainingPages[page] = true
				existingPages[page] = true
			}

			pageToPrecedingCount := make(map[int]int)
			for _, from := range update.pages {
				for _, to := range pageToDependentPages[from] {
					_, ok := existingPages[to]
					if ok {
						pageToPrecedingCount[to] += 1
					}
				}
			}

			orderedPages := []int{}
			for len(remainingPages) > 0 {
				for page := range remainingPages {
					if pageToPrecedingCount[page] == 0 {
						delete(remainingPages, page)
						for _, to := range pageToDependentPages[page] {
							_, ok := existingPages[to]
							if ok {
								pageToPrecedingCount[to] -= 1
							}
						}
						orderedPages = append(orderedPages, page)
					}
				}

			}
			middleIndex := len(orderedPages) / 2
			middlePagesSum += orderedPages[middleIndex]
		}
	}

	return middlePagesSum
}
