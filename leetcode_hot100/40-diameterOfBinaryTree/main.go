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
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	var getDepth func(root *TreeNode)
	var resp int
	getDepth = func(root *TreeNode) {
		if root != nil {

			getDepth(root.Right)
			m := maxDepth(root.Right) + maxDepth(root.Left)
			if resp < m {
				resp = m
			}

			getDepth(root.Left)

		}
	}
	getDepth(root)

	return resp
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
