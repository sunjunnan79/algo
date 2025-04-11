package main

func main() {

}
func rightSideView(root *TreeNode) []int {
	var f func(root *TreeNode, temp int)
	var res []int

	f = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}
		if len(res) <= temp {
			res = append(res, 0)
		}
		res[temp] = root.Val
		f(root.Left, temp+1)
		f(root.Right, temp+1)
	}
	f(root, 0)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
