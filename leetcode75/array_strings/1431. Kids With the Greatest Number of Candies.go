package array_strings

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := 0
	for i := 0; i < len(candies); i++ {
		if greatest < candies[i] {
			greatest = candies[i]
		}
	}

	var results []bool
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatest {
			results = append(results, true)
		} else {
			results = append(results, false)
		}
	}

	return results
}
