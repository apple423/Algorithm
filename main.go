package main

import (
	"fmt"
	"leetcode/75/day_5"
)

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	// nums := []int{7, 6, 5, 4, 3, 1}
	// nums := []int{2, 4, 1}
	// results := day_one.PivotIndex(nums)
	// a := &day_3.ListNode{Val: 1, Next: &day_3.ListNode{Val: 2, Next: &day_3.ListNode{Val: 4, Next: nil}}}
	// a := &day_three.ListNode{}
	// b := &day_three.ListNode{Val: 1, Next: &day_three.ListNode{Val: 3, Next: &day_three.ListNode{Val: 4, Next: nil}}}

	result := day_5.MaxProfit(nums)
	// for result != nil {
	// 	fmt.Print(result.Val)
	// 	result = result.Next
	// }
	fmt.Printf("results: %+v\n", result)
}
