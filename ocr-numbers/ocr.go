package ocr

import (
	"fmt"
	"strings"
	"bytes"
)

var digitMap = map[string]byte {
	`
 _ 
| |
|_|
   
`:'0',
   `
   
  |
  |
   
`:'1',
   `
 _ 
 _|
|_ 
   
`:'2',
   `
 _ 
 _|
 _|
   
`:'3',
   `
   
|_|
  |
   
`:'4',
   `
 _ 
|_ 
 _|
   
`:'5',
   `
 _ 
|_ 
|_|
   
`:'6',
   `
 _ 
  |
  |
   
`:'7',
   `
 _ 
|_|
|_|
   
`:'8',
   `
 _ 
|_|
 _|
   
`:'9',
}

func Recognize(input string) []string {
	blockLines := blockLinesFrom(input)
	output := []string{}
	for _, line := range blockLines {
		output = append(output, recognizeLine(line))
	}
	return output
}

func recognizeLine(line []string) string {
	var buffer bytes.Buffer
	for _, block := range line {
		buffer.WriteByte(recognizeDigit(block))
	}
	return buffer.String()
}

func recognizeDigit(input string) byte {
	if digit, ok := digitMap[input]; ok {
		return digit
	}
	return '?'
}

func blockLinesFrom(input string) [][]string {
	lines := strings.Split(input, "\n")[1:]
	numberOfLines := len(lines) / 4
	blocks := [][]string{}
	for j := 0; j < numberOfLines; j++ {
		blocks = append(blocks, blocksFrom(lines[4 * j: 4 * j + 4]))
	}
	return blocks
}

func blocksFrom(lines []string) []string {
	blocks := []string{}
	numberOfBlocks := len(lines[0]) / 3
	for i := 0; i < numberOfBlocks; i++ {
		var buffer bytes.Buffer
		buffer.WriteByte('\n')
		for j := 0; j < 4; j++ {
			buffer.WriteString(fmt.Sprintln(lines[j][3 * i:3 * i + 3]))
		}
		blocks = append(blocks, buffer.String())
	}
	return blocks
}