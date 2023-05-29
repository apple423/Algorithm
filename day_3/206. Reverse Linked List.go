package day_3

func ReverseList(head *ListNode) *ListNode {
	// iterate
	// tail := &ListNode{}
	// var arr []int
	// for head != nil && head.Val != 0 {
	// 	arr = append(arr, head.Val)
	// 	head = head.Next
	// }

	// if len(arr) == 0 {
	// 	return nil
	// }

	// tmp := tail
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	tmp.Val = arr[i]
	// 	if i != 0 {
	// 		tmp.Next = &ListNode{}
	// 	}
	// 	tmp = tmp.Next
	// }
	// return tail

	// iterate2
	// var revHead *ListNode
	// for head != nil {
	// 	tmp := head.Next
	// 	head.Next = revHead
	// 	revHead = head
	// 	head = tmp
	// }

	// return revHead

	return recursiveExam(head)
}

func recursiveExam(head *ListNode) *ListNode {
	revHead := head

	if head != nil && head.Next != nil {
		revHead = ReverseList(head.Next)
		head.Next.Next = head
		head.Next = nil
	}

	return revHead
}
