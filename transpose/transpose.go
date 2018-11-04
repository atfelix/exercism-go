package transpose

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}
	totalNumberOfRows := maxRows(input)
	totalNumberOfColumns := len(input)

	output := make([]string, totalNumberOfRows[0])
	for column := 0; column < totalNumberOfColumns; column++ {
		for row := 0; row < totalNumberOfRows[column]; row++ {
			if row < len(input[column]) {
				output[row] += string(input[column][row])
			} else {
				output[row] += " "
			}
		}
	}
	return output
}

func maxRows(input []string) []int {
	maxForEachColumn := []int{}
	for _ = range input {
		maxForEachColumn = append(maxForEachColumn, 0)
	}
	maxForEachColumn[len(input) - 1] = len(input[len(input) - 1])
	for i := len(input) - 2; i >= 0; i-- {
		maxForEachColumn[i] = max(maxForEachColumn[i + 1], len(input[i]))
	}
	return maxForEachColumn
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}