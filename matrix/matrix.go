package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix struct{
	numRows, numColumns int
	entries []int
}

func New(input string) (*matrix, error) {
	if !isMatrixValid(input) {
		return nil, errors.New("Invalid input for matrix")
	}
	parsedData := parse(input)
	m := &matrix{}
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

func (m matrix) Rows() [][]int {
	rows := [][]int{}
	for index := 0; index < m.numRows; index++ {
		rows = append(rows, m.row(index))
	}
	return rows
}

func (m matrix) row(index int) []int {
	row := []int{}
	for column := 0; column < m.numColumns; column++ {
		row = append(row, m.entryAt(index, column))
	}
	return row
}

func (m matrix) entryAt(row, column int) int {
	return m.entries[row * m.numColumns + column]
}

func (m matrix) Cols() [][]int {
	columns := [][]int{}
	for index := 0; index < m.numColumns; index++ {
		columns = append(columns, m.column(index))
	}
	return columns
}

func (m matrix) column(index int) []int {
	columns := []int{}
	for row := 0; row < m.numRows; row++ {
		columns = append(columns, m.entryAt(row, index))
	}
	return columns
}

func (m *matrix) Set(row, column, value int) bool {
	if !(m.isValid(row, column)) {
		return false
	}
	m.setEntryAt(row, column, value)
	return true
}

func (m *matrix) isValid(row, column int) bool {
	return 0 <= row && row < m.numRows && 0 <= column && column < m.numColumns
}

func (m *matrix) setEntryAt(row, column, value int) {
	m.entries[row * m.numColumns + column] = value
}