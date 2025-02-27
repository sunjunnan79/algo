package main

import "fmt"

func main() {
	fmt.Println(spiralOrder(
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}))

}

func spiralOrder(matrix [][]int) []int {
	var res []int

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res // 如果矩阵为空，直接返回空切片
	}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// 从左到右遍历上边界
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上到下遍历右边界
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			// 从右到左遍历下边界
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// 从下到上遍历左边界
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}
