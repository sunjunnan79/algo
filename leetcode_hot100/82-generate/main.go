package main

import "fmt"

func main() {
	fmt.Println(generate(3))
}

func generate(numRows int) [][]int {
	//维护一个最大长度为numRows-1的数组
	var result = make([][]int, numRows)

	//初始化空间并全部置为1
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			result[i][j] = 1
		}
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
