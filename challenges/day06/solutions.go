package main

type PositionWithDirection struct {
	row       int
	col       int
	direction int
}

type Position struct {
	row int
	col int
}

const (
	up = iota
	right
	down
	left
)

func findGuard(area [][]rune) PositionWithDirection {
	guard := PositionWithDirection{}
	for row, rowData := range area {
		for col, symbol := range rowData {
			if symbol == '^' {
				guard.row = row
				guard.col = col
				guard.direction = up
				return guard
			}
		}
	}
	return guard
}

func moveGuard(area [][]rune, guard PositionWithDirection, extraObstruction Position) PositionWithDirection {
	newRow := guard.row
	newCol := guard.col

	switch guard.direction {
	case up:
		newRow -= 1
	case right:
		newCol += 1
	case down:
		newRow += 1
	case left:
		newCol -= 1
	}

	newDirection := guard.direction
	if area[newRow][newCol] == '#' || (extraObstruction.row == newRow && extraObstruction.col == newCol) {
		newRow = guard.row
		newCol = guard.col
		switch guard.direction {
		case up:
			newDirection = right
		case right:
			newDirection = down
		case down:
			newDirection = left
		case left:
			newDirection = up
		}
	}
	return PositionWithDirection{row: newRow, col: newCol, direction: newDirection}
}

func containsLoop(area [][]rune, guard PositionWithDirection, extraObstruction Position) bool {
	currentlyVisited := map[PositionWithDirection]bool{guard: true}
	for guard.row > 0 && guard.row < len(area)-1 && guard.col > 0 && guard.col < len(area[0])-1 {
		guard = moveGuard(area, guard, extraObstruction)
		_, ok := currentlyVisited[guard]
		if ok {
			return true
		}
		currentlyVisited[guard] = true
	}
	return false
}

func solveFirst(area [][]rune) int {
	guard := findGuard(area)

	visited := map[Position]bool{{row: guard.row, col: guard.col}: true}

	for guard.row > 0 && guard.row < len(area)-1 && guard.col > 0 && guard.col < len(area[0])-1 {
		guard = moveGuard(area, guard, Position{row: -1, col: -1})
		visited[Position{row: guard.row, col: guard.col}] = true
	}

	return len(visited)
}

func solveSecond(area [][]rune) int {
	initialGuard := findGuard(area)

	originallyVisited := map[Position]bool{}
	loopingObstructions := 0

	guard := initialGuard
	for guard.row > 0 && guard.row < len(area)-1 && guard.col > 0 && guard.col < len(area[0])-1 {
		guard = moveGuard(area, guard, Position{row: -1, col: -1})
		originallyVisited[Position{row: guard.row, col: guard.col}] = true
	}

	for extraObstruction := range originallyVisited {
		if containsLoop(area, initialGuard, extraObstruction) {
			loopingObstructions += 1
		}
	}

	return loopingObstructions
}
