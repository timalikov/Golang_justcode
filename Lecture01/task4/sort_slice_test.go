package task4

import (
	"reflect"
	"testing"
)

func SortSliceTest(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 2, 3, 1}, []int{1, 2, 3, 4}},
		{[]int{5, 3, 8, 6, 2}, []int{2, 3, 5, 6, 8}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		SortSlice(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("bubbleSort(%v) = %v; expected %v", test.input, test.input, test.expected)
		}
	}

}
