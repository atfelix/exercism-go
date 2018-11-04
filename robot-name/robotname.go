package robotname

import (
	"math/rand"
)

type Robot struct {
	name string
}

var nameRegistry = map[string]bool{}
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "1234567890"

func (robot *Robot) Name() string {
	for robot.name == "" {
		name := randomTwoLetter() + randomThreeDigits()
		if !nameRegistry[name] {
			robot.name = name
			nameRegistry[name] = true
		}
	}
	return robot.name
}

func (robot *Robot) Reset() {
	robot.name = ""
}

func randomTwoLetter() string {
	return randomString(letters, 2)
}

func randomThreeDigits() string {
	return randomString(digits, 3)
}

func randomString(fromString string, ofLength int) string {
	bytes := []byte{}
	for index := 0; index < ofLength; index++ {
		bytes = append(bytes, fromString[rand.Intn(len(fromString))])
	}
	return string(bytes)
}