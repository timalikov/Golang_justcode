package task2

import (
	"reflect"
	"testing"
)

func LongestPrefixTest(t *testing.T) {
	tests := []struct {
		strs   []string
		output string
	}{
		{strs: []string{"flower", "flow", "flight"}, output: "fl"},
		{strs: []string{"dog", "racecar", "car"}, output: ""},
	}

	for _, tt := range tests {
		result := LongestCommonPrefix(tt.strs)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("LongestCommonPrefix(%v) = %v; want %v", tt.strs, result, tt.output)
		}
	}
}
