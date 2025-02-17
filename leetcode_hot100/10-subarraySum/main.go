package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))

}

func subarraySum(nums []int, k int) int {
	var resp int
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			if sum+nums[j] == k {
				resp++
			}
			sum += nums[j]
		}
	}
	return resp
}
