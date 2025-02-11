package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)

}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
