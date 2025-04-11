package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
}

func maxProduct(nums []int) int {
	maxVal, minVal := nums[0], nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		cur := nums[i]

		// 由于 maxVal 和 minVal 会被更新，先暂存
		maxTmp := max(cur, cur*maxVal, cur*minVal)
		minTmp := min(cur, cur*maxVal, cur*minVal)

		maxVal = maxTmp
		minVal = minTmp

		res = max(res, maxVal)
	}

	return res
}
