package main

func main() {
	sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func sortedArrayToBST(nums []int) *TreeNode {
	var getTree func(list []int) *TreeNode

	getTree = func(list []int) *TreeNode {
		if len(list) == 0 {
			return nil
		}

		var half = len(list) / 2

		root := &TreeNode{
			Val: list[half],
		}

		root.Left = getTree(list[:half])
		root.Right = getTree(list[half+1:])
		return root
	}

	return getTree(nums)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
