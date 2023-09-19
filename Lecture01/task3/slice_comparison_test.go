package task3

import (
	"reflect"
	"testing"
)

func SliceComparisonTest(t *testing.T) {
	tests := []struct {
		slice1 []int
		slice2 []int
		output bool
	}{
		{slice1: []int{1, 2, 3, 4, 5, 6}, slice2: []int{6, 5, 4, 3, 2, 1}, output: true},
		{slice1: []int{6, 7, 4, 5}, slice2: []int{4, 5, 3, 1, 8}, output: false},
		{[]int{9, 3, 6335, 2}, []int{2, 9, 6335, 3}, true},
	}

	for _, tt := range tests {
		result := SliceComparison(tt.slice1, tt.slice2)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("SliceComparison(%v, %b) = %v; want %v", tt.slice1, tt.slice2, result, tt.output)
		}
	}

}
