package accumulate

func Accumulate(list []string, converter func(string) string) []string {
	accumulation := []string{}
	for _, s := range list {
		accumulation = append(accumulation, converter(s))
	}
	return accumulation
}