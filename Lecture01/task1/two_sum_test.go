package task1

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, output: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, output: []int{1, 2}},
	}

	for _, tt := range tests {
		result := TwoSum(tt.nums, tt.target)
		if !slicesEqualUnordered(result, tt.output) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, result, tt.output)
		}
	}
}

func slicesEqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[int]int)
	for _, num := range a {
		count[num]++
	}
	for _, num := range b {
		count[num]--
		if count[num] < 0 {
			return false
		}
	}
	return true
}
