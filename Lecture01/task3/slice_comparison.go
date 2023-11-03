package task3

// Time complexity O(n + m)
func SliceComparison(firstSlice []int, secondSlice []int) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}

	elementCount := make(map[int]int)

	for _, el := range firstSlice {
		elementCount[el]++
	}

	for _, el := range secondSlice {
		if elementCount[el] == 0 {
			return false
		}
		elementCount[el]--
	}

	for _, count := range elementCount {
		if count != 0 {
			return false
		}
	}

	return true
}
