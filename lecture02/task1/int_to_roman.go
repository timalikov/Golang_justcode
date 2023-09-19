package task1

// Time complexity is: O(1)
func IntToRoman(num int) string {
	valSymbols := []struct {
		val    int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	romanNumeral := ""

	for _, pair := range valSymbols {
		for num >= pair.val {
			romanNumeral += pair.symbol
			num -= pair.val
		}
	}

	return romanNumeral
}
