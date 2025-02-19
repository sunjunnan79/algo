package main

import "fmt"

func main() {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	var symmetric func(left *TreeNode, right *TreeNode) bool

	symmetric = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)

	}

	return symmetric(root.Left, root.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
