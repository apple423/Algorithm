package array_strings

import "strings"

// not yet
func GcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		if strings.Contains(str1, str2) {
			result := str1[0 : len(str1)-len(str2)]
			if result == str1[len(str2):] {
				return result
			}
		}
	} else {
		if strings.Contains(str2, str1) {
			result := str2[0 : len(str2)-len(str1)]
			if result == str2[len(str1):] {
				return result
			}
		}
	}
	return ""
}
