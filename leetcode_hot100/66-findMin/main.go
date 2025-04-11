package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 8, 0, 1, 2}))
}

func findMin(nums []int) int {
	var left, right = 0, len(nums) - 1

	for left < right {

		mid := left + (right-left)/2

		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return nums[left]
}
