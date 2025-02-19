package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	var a, b, c = 0, 0, 1

	for i := 0; i < n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}
