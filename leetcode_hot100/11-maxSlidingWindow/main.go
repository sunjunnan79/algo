package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	var resp = make([]int, len(nums)-k+1)
	var m = 1
	//初始化阶段
	resp[0] = nums[0]

	for i := 1; i < k; i++ {
		//记录第一次的最大值
		resp[0] = max(resp[0], nums[i])
		//当前值大于最大值时则更新下标
		if max(nums[m], nums[i]) == nums[i] {
			m = i
		}
	}

	for i := 1; i <= len(nums)-k; i++ {

		if nums[m] < nums[i+k-1] {
			//如果新增值大于之前记录的值,则记录最大值为新增值并存储
			m = i + k - 1
			resp[i] = nums[m]
		} else {
			//如果并不大于之前的值,应当存储值为上一次存储的m,并更新m为新值(因为下一轮的默认已知最大值是该值)
			resp[i] = nums[m]
			if m > i {
				continue
			} else {
				//初始化值并尝试获取
				m = i + 1
				for j := i + 1; j < i+k; j++ {
					if max(nums[m], nums[j]) == nums[j] {
						m = j
					}
				}
			}

		}

	}
	return resp
}
