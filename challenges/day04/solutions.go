package main

func solveFirst(puzzle [][]rune) int {
	foundCount := 0
	searchWord := []rune("XMAS")

	for row, puzzleRow := range puzzle {
		for col := range puzzleRow {
			for verticalDir := -1; verticalDir <= 1; verticalDir++ {
				for horizontalDir := -1; horizontalDir <= 1; horizontalDir++ {
					if horizontalDir != 0 || verticalDir != 0 {
						new_row := row
						new_col := col
						searchWordId := 0
						for searchWordId < len(searchWord) && new_row >= 0 && new_row < len(puzzle) && new_col >= 0 && new_col < len(puzzleRow) && puzzle[new_row][new_col] == searchWord[searchWordId] {
							new_row += verticalDir
							new_col += horizontalDir
							searchWordId += 1
							if searchWordId == len(searchWord) {
								foundCount += 1
							}
						}
					}
				}
			}
		}
	}

	return foundCount
}

func solveSecond(puzzle [][]rune) int {
	foundCount := 0
	middleLetter := 'A'
	edgeLetterOne := 'M'
	edgeLetterTwo := 'S'

	for row, puzzleRow := range puzzle {
		for col := range puzzleRow {
			if puzzle[row][col] == middleLetter && row >= 1 && row < len(puzzle)-1 && col >= 1 && col < len(puzzleRow)-1 {
				if (puzzle[row-1][col-1] == edgeLetterOne && puzzle[row+1][col+1] == edgeLetterTwo) || (puzzle[row-1][col-1] == edgeLetterTwo && puzzle[row+1][col+1] == edgeLetterOne) {
					if (puzzle[row+1][col-1] == edgeLetterOne && puzzle[row-1][col+1] == edgeLetterTwo) || (puzzle[row+1][col-1] == edgeLetterTwo && puzzle[row-1][col+1] == edgeLetterOne) {
						foundCount += 1
					}
				}
			}
		}
	}

	return foundCount
}
