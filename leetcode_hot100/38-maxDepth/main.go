package main

import "fmt"

func main() {
	var root = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -1,
			Right: &TreeNode{
				Val: -2,
				Left: &TreeNode{
					Val: -3,
				},
			},
		},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {

	var depth func(root *TreeNode) int

	depth = func(root *TreeNode) int {

		if root != nil {
			return max(depth(root.Left), depth(root.Right)) + 1
		} else {
			return 0
		}
	}

	return depth(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
