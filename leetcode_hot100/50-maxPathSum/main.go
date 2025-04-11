package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
	}
	fmt.Println(maxPathSum(root))
}

//O(n^3)

//	func maxPathSum(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//		//获取树列表
//		var list []*TreeNode
//		var getlist func(root *TreeNode)
//		getlist = func(root *TreeNode) {
//			if root == nil {
//				return
//			}
//			list = append(list, root)
//			//遍历
//			getlist(root.Left)
//			getlist(root.Right)
//
//		}
//		getlist(root)
//
//		//获取路径函数
//		var getPath func(root *TreeNode, target *TreeNode, ls []*TreeNode) []*TreeNode
//		getPath = func(root *TreeNode, target *TreeNode, ls []*TreeNode) []*TreeNode {
//			if root == nil {
//				return nil
//			}
//
//			if root == target {
//				return append(ls, root)
//			}
//
//			l := getPath(root.Left, target, append(ls, root))
//			if l == nil {
//				l = getPath(root.Right, target, append(ls, root))
//			}
//
//			return l
//		}
//		//获取距离函数
//		var m int
//		var flag = true
//
//		var getRemote func(r1 *TreeNode, r2 *TreeNode)
//		getRemote = func(r1 *TreeNode, r2 *TreeNode) {
//			if r1 == nil {
//				return
//			}
//
//			//获取两个列表
//			l1 := getPath(root, r1, []*TreeNode{})
//			l2 := getPath(root, r2, []*TreeNode{})
//			//获取最后一个交叉节点
//			var n int
//			for i := 0; i < min(len(l1), len(l2)); i++ {
//				if l1[i] == l2[i] {
//					n = i
//				}
//			}
//			//获取距离
//			temp := l1[n].Val
//			for i := n + 1; i < len(l1); i++ {
//				temp += l1[i].Val
//			}
//
//			for i := n + 1; i < len(l2); i++ {
//				temp += l2[i].Val
//			}
//			if flag {
//				m = temp
//				flag = false
//			} else {
//				//比较距离大小
//				m = max(m, temp)
//			}
//
//		}
//
//		//遍历获取所有的距离
//		for i := 0; i < len(list); i++ {
//			for j := i; j < len(list); j++ {
//				getRemote(list[i], list[j])
//			}
//		}
//
//		return m
//	}
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxSum := -1 << 31

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))

		maxSum = max(maxSum, node.Val+left+right)

		return node.Val + max(left, right)
	}

	dfs(root)
	return maxSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
