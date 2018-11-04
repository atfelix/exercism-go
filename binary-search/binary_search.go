package binarysearch

func SearchInts(slice []int, key int) int {
	if slice == nil || len(slice) == 0 {
		return -1
	}
	return binarySearchHelper(slice, 0, len(slice) - 1, key)
}

func binarySearchHelper(slice []int, low, high, key int) int {
	for low < high {
		mid := (high - low) / 2 + low
		switch {
		case slice[mid] < key:
			low = mid + 1
		default:
			high = mid
		}
	}

	if slice[low] == key {
		return low
	}
	return -1
}