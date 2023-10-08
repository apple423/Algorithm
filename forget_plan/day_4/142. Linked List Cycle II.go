package day_4

func DetectCycle(head *ListNode) *ListNode {
	result := make(map[*ListNode]*ListNode)
	for head != nil {
		if result[head] != nil {
			return result[head]
		}
		result[head] = head
		head = head.Next
	}

	return nil
}

// Two pointer
// func detectCycle(head *ListNode) *ListNode {
// 	slow := head
// 	fast := head

// 	for fast != nil && fast.Next != nil {
// 		slow = slow.Next
// 		fast = fast.Next.Next

// 		if slow == fast {
// 			slow = head
// 			for slow != fast {
// 				slow = slow.Next
// 				fast = fast.Next
// 			}
// 			return slow
// 		}
// 	}

// 	return nil
// }
