package day_5

func LongestPalindrome(s string) int {
	chars := make(map[string]int)
	for i := 0; i < len(s); i++ {
		chars[string(s[i])] = chars[string(s[i])] + 1
	}

	result := 0
	hasOdd := false
	for _, v := range chars {
		if v%2 == 0 {
			result = result + v
		} else {
			result = result + v - 1
			hasOdd = true
		}
	}

	if hasOdd {
		return result + 1
	}

	return result
}
