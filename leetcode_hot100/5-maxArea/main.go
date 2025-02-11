package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

func maxArea(height []int) int {
	result := 0
	i := 0
	j := len(height) - 1
	for i != j {
		size := min(height[i], height[j]) * (j - i)

		if size > result {
			result = size
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}
