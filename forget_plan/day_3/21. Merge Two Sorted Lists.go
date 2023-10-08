package day_3

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// recursive
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = MergeTwoLists(list2.Next, list1)
	return list2

	// iterate
	// result := &ListNode{}
	// tail := result

	// for list1 != nil && list2 != nil {
	// 	if list1.Val < list2.Val {
	// 		tail.Next = list1
	// 		list1 = list1.Next
	// 	} else {
	// 		tail.Next = list2
	// 		list2 = list2.Next
	// 	}
	// 	tail = tail.Next
	// }

	// if list1 == nil {
	// 	tail.Next = list2
	// } else {
	// 	tail.Next = list1
	// }

	// return result.Next
}
