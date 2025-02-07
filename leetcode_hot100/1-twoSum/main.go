package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// 被gpt爆了,活用map存储补数实现了O(n)级别
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		// 计算补数
		remain := target - num
		// 判断补数是否已经在hashMap中出现过
		if j, exists := hashMap[remain]; exists {
			return []int{j, i}
		}
		// 如果没找到，则将当前数字及其下标加入hashMap中
		hashMap[num] = i
	}
	return []int{}
}
