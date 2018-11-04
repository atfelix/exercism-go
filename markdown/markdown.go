package markdown

// implementation to refactor

import (
	"regexp"
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	return convertLines(markdown)
}

var (
	reHeader = regexp.MustCompile("#+ ")
	reList = regexp.MustCompile("(?m)(<li>.*</li>[\\n]?)+")
	reDouble = regexp.MustCompile("__([^_]*)__")
	reSingle = regexp.MustCompile("_([^_]*)_")
	doubleTemplate = "<strong>${1}</strong>"
	singleTemplate = "<em>${1}</em>"
	paragraphTemplate = "<p>%s</p>"
	listItemTemplate = "<li>%s</li>"
	listTemplate = "<ul>${0}</ul>\n"
	headerTemplate = "<h%d>%s</h%d>"
)


func convertLines(markdown string) string {
	lines := strings.Split(markdown, "\n")
	convertedLines := make([]string, len(lines))
	for index, line := range lines {
		match := reHeader.FindString(line)
		switch {
		case match != "": convertedLines[index] = convertHeader(line, len(match))
		case strings.HasPrefix(line, "* "): convertedLines[index] = convertListItem(line)
		default: convertedLines[index] = convertParagraph(line)
		}
	}
	html := convertList(strings.Join(convertedLines, "\n"))
	return strings.Replace(html, "\n", "", -1)
}

func convertHeader(line string, length int) string {
	return fmt.Sprintf(headerTemplate, length - 1, convert(line[length:]), length - 1)
}

func convertList(markdown string) string {
	return reList.ReplaceAllString(markdown, listTemplate)
}

func convertListItem(line string) string {
	return fmt.Sprintf(listItemTemplate, convert(line[2:]))
}

func convertParagraph(line string) string {
	return fmt.Sprintf(paragraphTemplate, convert(line))
}

func convertDoubleUnderscores(line string) string {
	return reDouble.ReplaceAllString(line, doubleTemplate)
}

func convertSingleUnderscores(line string) string {
	return reSingle.ReplaceAllString(line, singleTemplate)
}

func convert(line string) string {
	return convertSingleUnderscores(convertDoubleUnderscores(line))
}