package main

func main() {

}

func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}

	return dfs(node.Left, min, &node.Val) && dfs(node.Right, &node.Val, max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
