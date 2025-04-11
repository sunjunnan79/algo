package main

import "fmt"

func main() {

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

}

func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if res[0] != -1 && res[1] != -1 {
			return res
		}

		if res[0] == -1 && nums[i] == target {
			res[0] = i
		}

		if res[1] == -1 && nums[len(nums)-1-i] == target {
			res[1] = len(nums) - 1 - i
		}
	}
	return res
}
