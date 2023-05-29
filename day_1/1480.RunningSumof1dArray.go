package day_1

func RunningSum(nums []int) []int {
	var results = make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			results[i] = v
		} else {
			results[i] = results[i-1] + v
		}
	}

	return results
}
