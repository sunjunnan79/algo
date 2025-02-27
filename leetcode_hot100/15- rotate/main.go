package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(num, 3)
	fmt.Println(num)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	newNums := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(newNums); i++ {
		nums[i] = newNums[i]
	}

}
