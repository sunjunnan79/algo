package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	var m int
	for i := 0; i < len(nums); i++ {
		if m >= len(nums)-1 {
			return true
		}

		if i+nums[i] > m && i <= m {
			m = i + nums[i]
		}
	}
	return false
}

// gpt优化版本
//func canJump(nums []int) bool {
//    maxReach := 0 // 当前能跳到的最远位置
//    for i := 0; i <= maxReach; i++ {
//        maxReach = max(maxReach, i+nums[i]) // 更新最远可达位置
//        if maxReach >= len(nums)-1 {       // 如果最远可达位置已经覆盖终点，直接返回 true
//            return true
//        }
//    }
//    return false // 如果退出循环，说明无法到达终点
//}
