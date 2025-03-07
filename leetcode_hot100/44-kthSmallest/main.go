package main

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	var rangeTree func(root *TreeNode)
	var resp []int
	rangeTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		rangeTree(root.Left)
		resp = append(resp, root.Val)

		rangeTree(root.Right)
	}
	rangeTree(root)
	return resp[k-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
