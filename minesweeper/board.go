package minesweeper

import (
	"unicode"
	"fmt"
	"bytes"
)

type Board [][]byte

type entry [2]int

func (b Board) Count() error {
	if !b.isValid() {
		return fmt.Errorf("asdf")
	}
	for rowIndex, row := range b {
		for columnIndex := range row {
			if b[rowIndex][columnIndex] != ' ' {
				continue
			} else if count := b.minesAdjacentTo(rowIndex, columnIndex); count > 0 {
				b[rowIndex][columnIndex] = '0' + byte(count)
			}
		}
	}
	return nil
}

func (b Board) isValid() bool {
	for rowIndex, row := range b {
		for columnIndex := range row {
			if !b.isCharacterValid(rowIndex, columnIndex) {
				return false
			}
		}
	}
	return true
}

func (b Board) isCharacterValid(rowIndex, columnIndex int) bool {
	char := b[rowIndex][columnIndex]
	switch {
	case 0 < rowIndex && rowIndex < len(b) - 1 && 0 < columnIndex && columnIndex < len(b[0]) - 1:
		return unicode.IsDigit(rune(char)) || char == ' ' || isMine(char)
	case rowIndex == 0 && columnIndex == 0:
		return isCorner(char)
	case rowIndex == 0 && columnIndex == len(b[0]) - 1:
		return isCorner(char)
	case rowIndex == len(b) - 1 && columnIndex == 0:
		return isCorner(char)
	case rowIndex == len(b) - 1 && columnIndex == len(b[0]) - 1:
		return isCorner(char)
	case rowIndex == 0 || rowIndex == len(b) - 1:
		return isHorizontalEdge(char)
	case columnIndex == 0 || columnIndex == len(b[0]) - 1:
		return isVerticalEdge(char)
	default:
		return true
	}
}

func isCorner(char byte) bool {
	return char == '+'
}

func isHorizontalEdge(char byte) bool {
	return char == '-'
}

func isVerticalEdge(char byte) bool {
	return char == '|'
}

func (b Board) minesAdjacentTo(rowIndex, columnIndex int) int {
	count := 0
	for _, neighbor := range b.neighbors(rowIndex, columnIndex) {
		if isMine(b[neighbor[0]][neighbor[1]]) {
			count++
		}
	}
	return count
}

func isMine(char byte) bool {
	return char == '*'
}

func (b Board) neighbors(rowIndex, columnIndex int) []entry {
	entries := []entry{}

	if rowIndex == 0 || rowIndex == len(b) - 1 || columnIndex == 0 || columnIndex == len(b[0]) - 1 {
		return entries
	}

	for _, dRow := range []int{-1, 0, 1} {
		for _, dColumn := range []int{-1, 0, 1} {
			if dRow == 0 && dColumn == 0 {
				continue
			}
			entries = append(entries, entry{rowIndex + dRow, columnIndex + dColumn})
		}
	}

	return entries
}

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}
