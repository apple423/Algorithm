package array_strings

func ProductExceptSelf(nums []int) []int {
	productAll := 1
	zeroCount := 0
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		productAll = productAll * nums[i]
	}

	if zeroCount > 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = productAll
		} else if zeroCount > 0 {
			result[i] = 0
		} else {
			result[i] = productAll / nums[i]
		}
	}
	return result
}
