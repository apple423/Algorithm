package main

import (
	"fmt"
	"leetcode/leetcode75/array_strings"
)

func main() {
	// nums := []int{7, 1, 5, 3, 6, 4}
	// nums := []int{7, 6, 5, 4, 3, 1}
	// nums := []int{2, 4, 1}
	// results := day_one.PivotIndex(nums)
	// a := &day_3.ListNode{Val: 1, Next: &day_3.ListNode{Val: 2, Next: &day_3.ListNode{Val: 4, Next: nil}}}
	// a := &day_three.ListNode{}
	// b := &day_three.ListNode{Val: 1, Next: &day_three.ListNode{Val: 3, Next: &day_three.ListNode{Val: 4, Next: nil}}}
	// s := "abccccdd"

	// arr := []int{1, 0, 0, 0, 0, 1}
	arr := []int{1, 2, 3, 4}
	result := array_strings.ProductExceptSelf(arr)
	// for result != nil {
	// 	fmt.Print(result.Val)
	// 	result = result.Next
	// }
	fmt.Printf("results: %+v\n", result)
}
