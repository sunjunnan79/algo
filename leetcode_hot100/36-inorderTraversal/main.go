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
	fmt.Println(inorderTraversal(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var inorder func(root *TreeNode)
	var res []int
	inorder = func(root *TreeNode) {
		if root != nil {
			inorder(root.Left)
			res = append(res, root.Val)
			inorder(root.Right)
		}
	}
	inorder(root)
	return res
}
