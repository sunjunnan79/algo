package main

func main() {

}

func flatten(root *TreeNode) {
	var list []*TreeNode
	var rangeTree func(node *TreeNode)
	rangeTree = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		rangeTree(node.Left)
		rangeTree(node.Right)
	}
	rangeTree(root)
	if len(list) > 0 {
		for i := 0; i < len(list)-1; i++ {
			list[i].Left = nil
			list[i].Right = list[i+1]
		}
		list[len(list)-1].Right = nil
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
