package main

import "fmt"

var matrix = [][]int{
	{0, 1, 2},
	{1, 0, 3},
	{2, 1, 4},
}

func main() {
	setZeroes(matrix)
	for i := range matrix {
		fmt.Println(matrix[i])
	}

}
func setZeroes(matrix [][]int) {

	// 找到所有的zero
	var zeros [][2]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, [2]int{i, j})
			}
		}
	}

	for _, zero := range zeros {
		//处理整个行
		for i := 0; i < len(matrix[zero[0]]); i++ {
			matrix[zero[0]][i] = 0
		}

		//处理整个列
		for i := 0; i < len(matrix); i++ {
			matrix[i][zero[1]] = 0
		}

	}
}
