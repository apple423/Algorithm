package array_strings

import "strings"

func ReverseWords(s string) string {
	trim := strings.TrimSpace(s)
	strs := strings.Split(trim, " ")

	var result string
	for i, _ := range strs {
		if strs[len(strs)-1-i] == "" {
			continue
		}
		result += strs[len(strs)-1-i] + " "
	}

	return strings.TrimSpace(result)
}
