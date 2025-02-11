package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

// 自己的复杂解法....

//func trap(height []int) int {
//	var n = len(height)
//	var fast int
//	var area int
//	var slow int
//	var sum int
//
//	var high int
//	for {
//
//		//如果找到储水区域
//		if height[fast] >= height[slow] {
//			//记录下一次的起始高度的位置
//			slow = fast
//			//存储sum值到area中并清空
//			area += sum
//			sum = 0
//			high = slow + 1
//		} else {
//			//添加到当前面积的总和值
//			sum += height[slow] - height[fast]
//
//		}
//
//		//无法存水
//		if fast == n-1 {
//			slow++
//
//			//查看是否存在存水空间
//			if high != slow {
//				for i := slow; i < high; i++ {
//					area += height[high] - height[i]
//				}
//				slow = high
//			}
//			fast = slow
//			sum = 0
//			if slow >= n-1 {
//				break
//			} else {
//				high = slow + 1
//			}
//		}
//
//		//快指针自增
//		fast++
//
//		if height[fast] >= height[high] {
//			high = fast
//		}
//
//	}
//
//	return area
//}
