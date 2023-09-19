package task2

// Time complexity is O(n * m), where n is the number of strings and m is the length of the shortest string
func LongestCommonPrefix(strs []string) string {
	res := ""

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	for i, ch := range strs[0] {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || rune(strs[j][i]) != ch {
				return res
			}

		}
		res += string(ch)

	}
	return res

}
