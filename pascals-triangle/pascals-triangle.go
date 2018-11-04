package pascal

func Triangle(n int) [][]int {
	cache := [][]int{{1}}
	for i := 1; i < n; i++ {
		cache = append(cache, []int{1})
		for j := 1; j < i; j++ {
			cache[i] = append(cache[i], cache[i - 1][j - 1] + cache[i - 1][j])
		}
		cache[i] = append(cache[i], 1)
	}
	return cache
}