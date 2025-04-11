package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4, 5}))
}

func permute(nums []int) [][]int {
	//结果集
	var res [][]int
	var used = make([]bool, len(nums))
	var traceback = func(nums []int, path []int, used []bool) {}
	traceback = func(nums []int, path []int, used []bool) {
		//达成某种条件则添加到结果集
		if len(path) == len(nums) {

			p := make([]int, len(nums))
			copy(p, path)
			res = append(res, p)
			return
		}

		for i := 0; i < len(nums); i++ {

			//如果已经被使用了则跳过当前的数(因为是不能够重复使用的)
			if used[i] {
				continue
			}

			//表示当前位已经被使用
			used[i] = true
			path = append(path, nums[i])
			//递归调用
			traceback(nums, path, used)

			//清除上次递归调用的记录,继续去遍历下一个结果
			path = path[:len(path)-1]
			used[i] = false

		}

	}

	traceback(nums, []int{}, used)

	return res
}
