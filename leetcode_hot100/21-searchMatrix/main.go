package main

import "fmt"

func main() {
	var a = [][]int{
		{5, 6, 10, 14},
		{6, 10, 13, 18},
		{10, 13, 18, 19},
	}
	fmt.Println(searchMatrix(a, 14))
}

func searchMatrix(matrix [][]int, target int) bool {

	left := [2]int{0, 0}
	right := [2]int{len(matrix) - 1, len(matrix[0]) - 1}
	middle := [2]int{left[0] + (right[0]-left[0])/2, left[1] + (right[1]-left[1])/2}
	var temp int
	var flag bool
	for compare(left, right) {
		temp = matrix[middle[0]][middle[1]]
		switch {
		case temp > target:
			if right[0] == middle[0] && right[1] == middle[1] {
				flag = true
			}
			right[0] = middle[0]
			right[1] = middle[1]
		case temp < target:
			if left[0] == middle[0] && left[1] == middle[1] {
				flag = true
			}
			left[0] = middle[0]
			left[1] = middle[1]
		default:
			return true
		}
		if flag {
			break
		}
		middle = [2]int{left[0] + (right[0]-left[0])/2, left[1] + (right[1]-left[1])/2}
	}
	var a, b = left[0]*len(matrix[0]) + left[1] + 1, right[0]*len(matrix[0]) + right[1] + 1
	for i := a; i <= b; i++ {
		if getData(matrix, i) == target {
			return true
		}
	}

	return false
}

func compare(left [2]int, right [2]int) bool {
	if left[0] == right[0] && left[1] == right[1] {
		return false
	} else if left[0] <= right[0] && left[1] <= right[1] {
		return true
	}
	return false
}

func getData(matrix [][]int, length int) int {
	row := length / len(matrix[0])
	col := length % len(matrix[0])
	if col == 0 {
		return matrix[row-1][len(matrix[0])-1]
	} else {
		return matrix[row][col-1]
	}
}
