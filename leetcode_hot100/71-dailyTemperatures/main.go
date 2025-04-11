package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{1, 6, 4, 8}))

}

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	var resp = make([]int, len(temperatures))
	for i := 0; i < len(temperatures); {
		if len(stack) == 0 || temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if temperatures[i] > temperatures[stack[len(stack)-1]] {
			//如果找到大于当前值的则stack缩减，并且不移动i使其一直尝试知道直到失败为止
			resp[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		}
		i++
	}
	return resp
}
