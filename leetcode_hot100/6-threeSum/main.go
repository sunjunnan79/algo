package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}

func threeSum(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	//sort.Ints(nums)

	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		// 跳过重复的 nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, n-1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				// 找到一组解
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的 left 和 right
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移动指针
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
