package task1

// Time complexity: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, exists := seen[complement]; exists {
			return []int{i, j}
		}

		seen[num] = i
	}

	return nil

}
