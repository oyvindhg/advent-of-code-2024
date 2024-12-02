package main

func isReportSafe(report []int, skipNumber int) (bool, int) {
	isSafe := true
	isIncreasing := true
	prevLimit := report[0]
	newId := 0
	failedId := -1
	for originalId, limit := range report {
		if originalId != skipNumber {
			if newId == 1 {
				if limit < prevLimit {
					isIncreasing = false
				}
			}
			if newId >= 1 {
				if isIncreasing {
					if limit > prevLimit+3 || limit < prevLimit+1 {
						isSafe = false
						if failedId == -1 {
							failedId = originalId
						}
					}
				} else {
					if limit < prevLimit-3 || limit > prevLimit-1 {
						isSafe = false
						if failedId == -1 {
							failedId = originalId
						}
					}
				}
			}
			prevLimit = limit
			newId += 1
		}
	}
	return isSafe, failedId
}

func solveFirst(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe, _ := isReportSafe(report, -1)
		if isSafe {
			safeReports += 1
		}
	}

	return safeReports
}

func solveSecond(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		isSafe, failedId := isReportSafe(report, -1)
		if !isSafe {
			isSafe, _ = isReportSafe(report, failedId)
			if !isSafe {
				isSafe, _ = isReportSafe(report, failedId-1)
			}
			if !isSafe && failedId == 2 {
				isSafe, _ = isReportSafe(report, 0)
			}
		}
		if isSafe {
			safeReports += 1
		}
	}

	return safeReports
}
