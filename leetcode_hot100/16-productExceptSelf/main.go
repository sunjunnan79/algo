package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))

}
func productExceptSelf(nums []int) []int {
	if len(nums) < 0 {
		return nil
	}

	var left, right []int
	left = make([]int, len(nums))
	right = make([]int, len(nums))
	//初始化
	left[0] = 1
	right[len(nums)-1] = 1
	// 遍历获得所有左侧值
	for i := 0; i < len(nums)-1; i++ {
		left[i+1] = left[i] * nums[i]
	}

	//同理遍历获取所有右侧值
	for i := len(nums) - 1; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	var resp []int
	for i := 0; i < len(nums); i++ {
		resp = append(resp, left[i]*right[i])
	}

	return resp
}
