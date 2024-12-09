package main

import (
	"math"
)

type OccupiedSpace struct {
	id    int // -1 means free space
	space int
}

func decompressDisk(diskmap []rune) []OccupiedSpace {
	decompressed := []OccupiedSpace{}
	for i, spaceRune := range diskmap {
		space := int(spaceRune - '0')
		if i%2 == 0 {
			decompressed = append(decompressed, OccupiedSpace{id: i / 2, space: space})
		} else {
			decompressed = append(decompressed, OccupiedSpace{id: -1, space: space})
		}
	}
	return decompressed
}

func calculateChecksum(moved []OccupiedSpace) int {
	position := 0
	checksum := 0
	for _, occupiedSpace := range moved {
		for i := position; i < position+occupiedSpace.space; i++ {
			if occupiedSpace.id != -1 {
				checksum += i * occupiedSpace.id
			}
		}
		position += occupiedSpace.space
	}

	return checksum
}

func solveFirst(diskmap []rune) int {
	decompressed := decompressDisk(diskmap)

	moved := []OccupiedSpace{}

	startIndex := 0
	endIndex := len(decompressed) - 1
	remainingAtEnd := decompressed[endIndex].space
	for startIndex <= endIndex {
		if startIndex == endIndex {
			moved = append(moved, OccupiedSpace{id: decompressed[endIndex].id, space: remainingAtEnd})
		} else if decompressed[startIndex].id != -1 {
			moved = append(moved, decompressed[startIndex])
		} else {
			free := decompressed[startIndex].space
			for free > 0 && endIndex > startIndex {
				using := int(math.Min(float64(free), float64(remainingAtEnd)))
				free -= using
				remainingAtEnd -= using
				moved = append(moved, OccupiedSpace{id: decompressed[endIndex].id, space: using})
				if remainingAtEnd == 0 {
					endIndex -= 2
					remainingAtEnd = decompressed[endIndex].space
				}
			}
		}
		startIndex += 1
	}

	return calculateChecksum(moved)
}

func solveSecond(diskmap []rune) int {
	decompressed := decompressDisk(diskmap)

	moved := []OccupiedSpace{}

	movedIndices := make(map[int]bool)
	for start, occupiedSpace := range decompressed {

		if occupiedSpace.id != -1 {
			_, ok := movedIndices[occupiedSpace.id]
			if !ok {
				moved = append(moved, occupiedSpace)
			} else {
				moved = append(moved, OccupiedSpace{id: -1, space: occupiedSpace.space})
			}
		} else {
			remaining := occupiedSpace.space
			for end := len(decompressed) - 1; end > start; end-- {
				_, ok := movedIndices[decompressed[end].id]
				if decompressed[end].id != -1 && !ok {
					if decompressed[end].space <= remaining {
						moved = append(moved, decompressed[end])
						remaining -= decompressed[end].space
						movedIndices[decompressed[end].id] = true
					}
				}
			}
			if remaining > 0 {
				moved = append(moved, OccupiedSpace{id: -1, space: remaining})
			}
		}
	}

	return calculateChecksum(moved)
}
