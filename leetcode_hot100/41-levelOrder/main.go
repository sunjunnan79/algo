package main

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	var resp [][]int
	var rangeTree func(root *TreeNode, temp int)
	rangeTree = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}

		if len(resp) <= temp {
			resp = append(resp, []int{})
		}

		rangeTree(root.Left, temp+1)
		resp[temp] = append(resp[temp], root.Val)
		rangeTree(root.Right, temp+1)

	}
	rangeTree(root, 0)
	return resp
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
