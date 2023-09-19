package task1

import (
	"reflect"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num    int
		output string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		result := IntToRoman(tt.num)
		if !reflect.DeepEqual(result, tt.output) {
			t.Errorf("IntToRoman(%v) = %v; want %v", tt.num, result, tt.output)
		}
	}
}
