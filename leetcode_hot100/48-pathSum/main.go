package main

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	var hash = make(map[int]int)
	//开始到0有一条路径
	hash[0] = 1
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val

		ans += hash[sum-targetSum]

		hash[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		hash[sum]--
	}

	dfs(root, 0)

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
