package array_strings

func MergeAlternately(word1 string, word2 string) string {
	var result string
	n := len(word1)
	m := len(word2)
	if n > m {
		for i := 0; i < m; i++ {
			result += string(word1[i])
			result += string(word2[i])
		}
		result += string(word1[m:])
	} else {
		for i := 0; i < n; i++ {
			result += string(word1[i])
			result += string(word2[i])
		}
		if n != m {
			result += string(word2[n:])
		}
	}

	return result
}
