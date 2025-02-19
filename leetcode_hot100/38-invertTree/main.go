package main

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
	invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(root *TreeNode)

	invert = func(root *TreeNode) {
		if root != nil {
			invert(root.Left)
			invert(root.Right)
			root.Right, root.Left = root.Left, root.Right
		}

	}

	invert(root)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
