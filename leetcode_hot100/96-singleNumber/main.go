package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2}))
}

func singleNumber(nums []int) int {
	resp := 0
	for _, num := range nums {
		resp ^= num
	}
	return resp
}

//无敌,我真想不到:
//算法核心思想：
//假设数组中的所有数字除了一个只出现一次外，其他数字都出现了两次。
//利用异或的特性：
//
//两个相同的数字异或结果为 0。
//把所有数字进行异或操作后，成对的数字会相互抵消（异或为 0），最终剩下的就是只出现一次的数字。
