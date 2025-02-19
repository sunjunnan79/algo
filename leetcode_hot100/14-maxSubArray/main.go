package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}

func maxSubArray(nums []int) int {
	//初始化
	var resp int
	if len(nums) != 0 {
		resp = nums[0]
	}

	//遍历的长度
	for j := 1; j <= len(nums); j++ {
		//初始化
		var temp int
		//遍历数组
		for i := 0; i < len(nums); i++ {
			temp += nums[i]
			if i >= j {
				temp -= nums[i-j]

			}
			resp = max(resp, temp)
		}

	}
	return resp
}
