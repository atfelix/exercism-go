package bob

import "strings"

func Hey(remark string) string {
	remark = strings.Trim(remark, " ")
	switch {
		case isSilence(remark): return "Fine. Be that way!"
		case isQuestion(remark) && isYelling(remark): return "Calm down, I know what I'm doing!"
		case isYelling(remark): return "Whoa, chill out!"
		case isQuestion(remark): return "Sure."
		default: return "Whatever."
	}
}

func isSilence(remark string) bool {
	return len(strings.Fields(remark)) == 0
}

func isYelling(remark string) bool {
	return hasAlpha(remark) && remark == strings.ToUpper(remark)
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func hasAlpha(remark string) bool {
	return strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvwxyz")
}
