package day_4

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	result := make(map[int]*ListNode)
	size := 0
	for head != nil {
		result[size] = head
		size++
		head = head.Next
	}

	mid := size / 2
	return result[mid]
}
