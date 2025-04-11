package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for {
			//如果不相等就尝试交换
			if nums[i] != i+1 {
				if nums[i] == nums[nums[i]-1] {
					return nums[i]
				} else {
					nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
				}
			} else {
				break
			}
		}

	}
	return 0
}
