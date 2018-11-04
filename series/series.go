package series

func All(length int, s string) []string {
	strings := []string{}
	for i := 0; i + length <= len(s); i++ {
		strings = append(strings, s[i:i + length])
	}
	return strings
}

func UnsafeFirst(length int, s string) (result string) {
	return s[:length]
}

func First(length int, s string) (string, bool) {
	if len(s) < length {
		return "", false
	}
	return UnsafeFirst(length, s), true
}