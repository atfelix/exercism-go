package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	numRows, numColumns int
	entries             []int
}

func New(input string) (*Matrix, error) {
	if !isMatrixValid(input) {
		return nil, errors.New("Invalid input for Matrix")
	}
	parsedData := parse(input)
	m := &Matrix{}
	m.numRows = len(parsedData)
	m.numColumns = len(strings.Split(parsedData[0], " "))

	for _, row := range parsedData {
		for _, entry := range strings.Split(row, " ") {
			integer, err := strconv.Atoi(entry)
			if err != nil {
				return nil, err
			}
			m.entries = append(m.entries, integer)
		}
	}
	return m, nil
}

func isMatrixValid(input string) bool {
	rows := strings.Split(input, "\n")
	lengthOfColumn := len(strings.Split(rows[0], " "))
	for _, row := range rows[1:] {
		row = strings.Trim(row, " ")
		if lengthOfColumn != len(strings.Split(row, " ")) {
			return false
		}
	}
	return true
}

func parse(input string) []string {
	slice := []string{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		slice = append(slice, strings.Trim(row, " "))
	}
	return slice
}

func (m Matrix) rows() [][]int {
	rows := [][]int{}
	for index := 0; index < m.numRows; index++ {
		rows = append(rows, m.row(index))
	}
	return rows
}

func (m Matrix) row(index int) []int {
	row := []int{}
	for column := 0; column < m.numColumns; column++ {
		row = append(row, m.entryAt(index, column))
	}
	return row
}

func (m Matrix) entryAt(row, column int) int {
	return m.entries[row*m.numColumns+column]
}

func (m Matrix) columns() [][]int {
	columns := [][]int{}
	for index := 0; index < m.numColumns; index++ {
		columns = append(columns, m.column(index))
	}
	return columns
}

func (m Matrix) column(index int) []int {
	columns := []int{}
	for row := 0; row < m.numRows; row++ {
		columns = append(columns, m.entryAt(row, index))
	}
	return columns
}

type Pair struct {
	row, column int
}

func (m Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for pair := range m.potentialSaddlePoints() {
		if pair.isSaddlePointIn(m) {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func (pair Pair) isSaddlePointIn(m Matrix) bool {
	row, column := m.row(pair.row), m.column(pair.column)
	maxInRow, minInColumn := max(row), min(column)
	entry := m.entryAt(pair.row, pair.column)
	return entry == maxInRow && entry == minInColumn
}

func (m Matrix) potentialSaddlePoints() map[Pair]bool {
	pairs := map[Pair]bool{}
	rows := m.rows()
	for index := range rows {
		for _, maxIndex := range maxIndices(m.row(index)) {
			pairs[Pair{row: index, column: maxIndex}] = true
		}
	}

	columns := m.columns()
	for index := range columns {
		for _, maxIndex := range maxIndices(m.column(index)) {
			pairs[Pair{row: maxIndex, column: index}] = true
		}
	}
	return pairs
}

func maxIndices(list []int) []int {
	maxElement := max(list)
	indices := []int{}
	for index := range list {
		if list[index] == maxElement {
			indices = append(indices, index)
		}
	}
	return indices
}

func max(list []int) int {
	result := list[0]
	for _, element := range list {
		if result < element {
			result = element
		}
	}
	return result
}

func min(list []int) int {
	result := list[0]
	for _, element := range list {
		if result > element {
			result = element
		}
	}
	return result
}
