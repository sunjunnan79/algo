package main

import "fmt"

func main() {
	var a = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(a)

	for _, v := range a {
		fmt.Println(v)
	}

}
func rotate(matrix [][]int) {
	//倒置
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//交换
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[j][i], matrix[j][len(matrix)-i-1] = matrix[j][len(matrix)-i-1], matrix[j][i]
		}
	}
}
