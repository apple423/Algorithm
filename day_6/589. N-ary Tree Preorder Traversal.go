package day_6

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	result = append(result, root.Val)

	if root == nil {
		return result
	}

	return result
}
