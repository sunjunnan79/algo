package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 5}))
}
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{}
	maxArea := 0

	// 右边界 + 计算面积
	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}
	return maxArea
}

//func largestRectangleArea(heights []int) int {
//
//	if len(heights) == 0 {
//		return 0
//	}
//
//	//获取最高值
//	high := heights[0]
//	//获取最低值
//	low := heights[0]
//	//初始化目标列表
//	var list []int
//	for i := 0; i < len(heights); i++ {
//		high = max(high, heights[i])
//		low = min(low, heights[i])
//		list = append(list, i)
//	}
//
//	var m = len(list) * low
//	for i := low; i <= high; i++ {
//
//		// 层序并遍历所有结果
//		var temp []int
//		var count int
//
//		for j := 0; j < len(list); j++ {
//			//存储下一层需要遍历的对象
//			if heights[list[j]] >= i+1 {
//				temp = append(temp, list[j])
//				if count == 0 || list[j] == list[j-1]+1 {
//					count++
//				} else {
//					m = max(m, (i+1)*count)
//					count = 1
//				}
//			} else {
//				m = max(m, (i+1)*count)
//				count = 0
//			}
//		}
//
//		m = max(m, (i+1)*count)
//
//		//更新新的遍历对象
//		list = temp
//	}
//	return m
//}
