package main

type Location struct {
	row int
	col int
}

func findAntinodes(cityMap [][]rune, allowHarmonics bool) int {
	antennas := make(map[rune][]Location)

	for row, rowData := range cityMap {
		for col, letter := range rowData {
			if letter != '.' {
				antennas[letter] = append(antennas[letter], Location{row: row, col: col})
			}
		}
	}

	antinodes := make(map[Location]bool)
	for _, locations := range antennas {
		for _, firstLocation := range locations {
			for _, secondLocation := range locations {
				if firstLocation != secondLocation {
					rowDiff := firstLocation.row - secondLocation.row
					colDiff := firstLocation.col - secondLocation.col

					firstAntinode := Location{row: firstLocation.row, col: firstLocation.col}
					for {
						firstAntinode.row -= rowDiff
						firstAntinode.col -= colDiff
						if !allowHarmonics {
							if firstAntinode == secondLocation {
								firstAntinode.row -= rowDiff
								firstAntinode.col -= colDiff
							}
						}

						if firstAntinode.row >= 0 && firstAntinode.row < len(cityMap) && firstAntinode.col >= 0 && firstAntinode.col < len(cityMap[0]) {
							antinodes[firstAntinode] = true
						} else {
							break
						}
						if !allowHarmonics {
							break
						}
					}

					secondAntinode := Location{row: firstLocation.row, col: firstLocation.col}
					for {
						secondAntinode.row += rowDiff
						secondAntinode.col += colDiff

						if !allowHarmonics {
							if secondAntinode == secondLocation {
								secondAntinode.row += rowDiff
								secondAntinode.col += colDiff
							}
						}

						if secondAntinode.row >= 0 && secondAntinode.row < len(cityMap) && secondAntinode.col >= 0 && secondAntinode.col < len(cityMap[0]) {
							antinodes[secondAntinode] = true
						} else {
							break
						}
						if !allowHarmonics {
							break
						}
					}
				}
			}
		}
	}

	return len(antinodes)
}

func solveFirst(cityMap [][]rune) int {
	return findAntinodes(cityMap, false)
}

func solveSecond(cityMap [][]rune) int {
	return findAntinodes(cityMap, true)
}
