package main

import (
	"fmt"
	"time"
)

func main() {
	sort([]int{4, 8, 6, 1})
	time.Sleep(10 * time.Second)
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		go func() {
			time.Sleep(time.Duration(nums[i]) * time.Second)
			fmt.Println(nums[i])
		}()
	}
}
