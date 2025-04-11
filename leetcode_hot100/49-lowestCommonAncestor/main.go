package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	p := root.Left
	q := root.Left.Right.Right
	fmt.Println(lowestCommonAncestor(root, p, q))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(node *TreeNode, list []*TreeNode, target *TreeNode) []*TreeNode

	dfs = func(node *TreeNode, list []*TreeNode, target *TreeNode) []*TreeNode {

		if node == nil {
			return nil
		} else if node == target {
			return append(list, node)
		}

		l := dfs(node.Left, append(list, node), target)
		if l == nil {
			l = dfs(node.Right, append(list, node), target)
		}
		return l
	}

	l1 := dfs(root, []*TreeNode{}, p)
	l2 := dfs(root, []*TreeNode{}, q)
	m := min(len(l1), len(l2))
	var resp *TreeNode
	for i := 0; i < m; i++ {
		if l1[i] == l2[i] {
			resp = l1[i]
		}
	}

	return resp
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
