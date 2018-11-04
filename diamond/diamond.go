package diamond

import (
	"errors"
	"strings"
)

const (
	start = 'A'
	end = 'Z'
)

func Gen(b byte) (string, error) {
	if b < start || end < b {
		return "", errors.New("Byte out of range")
	}
	numRows := int(2 * (b - start) + 1)
	slice := make([]string, numRows)
	mid := numRows / 2
	slice[mid] = paddedString(b, 1) + paddedString(' ', numRows - 2) + paddedString(b, 1) + "\n"
	for index := 1; index < numRows / 2; index++ {
		s := paddedString(' ', index) + paddedString(b - byte(index), 1) + paddedString(' ', numRows - 2 * index - 2) + paddedString(b - byte(index), 1) + paddedString(' ', index) + "\n"
		slice[mid - index] = s
		slice[mid + index] = s
	}
	slice[0] = paddedString(' ', mid) + "A" + paddedString(' ', mid) + "\n"
	slice[len(slice) - 1] = slice[0]
	return strings.Join(slice, ""), nil
}

func paddedString(b byte, length int) string {
	s := ""
	for index := 0; index < length; index++ {
		s += string(b)
	}
	return s
}
