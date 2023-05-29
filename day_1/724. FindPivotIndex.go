package day_1

func PivotIndex(nums []int) int {
	for i := range nums {
		leftSum := 0
		for j := i - 1; j >= 0; j-- {
			leftSum = leftSum + nums[j]
		}

		rightSum := 0
		for j := i + 1; j < len(nums); j++ {
			rightSum = rightSum + nums[j]
		}

		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
