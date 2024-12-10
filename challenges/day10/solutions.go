package main

type Location struct {
	row int
	col int
}

func findTrailScore(start Location, heightMap [][]int, countDistinctPaths bool) int {
	stack := []Location{start}
	visited := make(map[Location]bool)
	score := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if heightMap[current.row][current.col] == 9 {
			score += 1
		} else {
			for rowDiff := -1; rowDiff <= 1; rowDiff++ {
				for colDiff := -1; colDiff <= 1; colDiff++ {
					if (rowDiff != 0 || colDiff != 0) && (rowDiff == 0 || colDiff == 0) {
						next := Location{row: current.row + rowDiff, col: current.col + colDiff}
						if next.row >= 0 && next.row < len(heightMap) && next.col >= 0 && next.col < len(heightMap[0]) {
							_, ok := visited[next]
							if !ok && (heightMap[next.row][next.col] == (heightMap[current.row][current.col] + 1)) {
								stack = append(stack, next)
								if !countDistinctPaths {
									visited[next] = true
								}
							}
						}
					}
				}
			}
		}
	}

	return score
}

func solveFirst(heightMap [][]int) int {
	totalScore := 0
	for row, rowData := range heightMap {
		for col, height := range rowData {
			if height == 0 {
				score := findTrailScore(Location{row: row, col: col}, heightMap, false)
				totalScore += score
			}
		}
	}

	return totalScore
}

func solveSecond(heightMap [][]int) int {
	totalScore := 0
	for row, rowData := range heightMap {
		for col, height := range rowData {
			if height == 0 {
				score := findTrailScore(Location{row: row, col: col}, heightMap, true)
				totalScore += score
			}
		}
	}

	return totalScore
}
