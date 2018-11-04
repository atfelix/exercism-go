package scale

import (
	"strings"
)

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var tonicSharps = []string{"C", "G", "f#", "a", "A"}
var tonicFlats = []string{"F", "bb", "Eb", "g", "d", "Db"}

func Scale(tonic, interval string) []string {
	notes := flats
	if isSharp(tonic) {
		notes = sharps
	}
	tonic = noteFrom(tonic)
	index := startIndex(notes, tonic)
	scale := []string{}
	count := 0
	
	for {
		if contains(scale, notes[index]) {
			break
		}
		scale = append(scale, notes[index])
		stepSize := step(interval, count)
		index = (index + stepSize) % len(notes)
		count++
	}
	return scale
}

func isSharp(tonic string) bool {
	for _, sharp := range tonicSharps {
		if sharp == tonic {
			return true
		}
	}
	return false
}

func startIndex(slice []string, query string) int {
	for index, s := range slice {
		if s == query {
			return index
		}
	}
	return -1
}

func contains(slice []string, s string) bool {
	for _, t := range slice {
		if s == t {
			return true
		}
	}
	return false
}

var scaleLengthMap = map[byte]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

func step(interval string, index int) int {
	if interval == "" {
		return 1
	}
	return scaleLengthMap[interval[index]]
}

func noteFrom(tonic string) string {
	return strings.ToUpper(string(tonic[0])) + tonic[1:]
}