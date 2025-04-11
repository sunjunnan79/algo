package main

import "fmt"

func main() {
	nums := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(nums))
}

func numIslands(grid [][]byte) int {
	var resp int

	//遍历所有节点,感觉不是很优秀
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '1' {
				dfs(grid, i, j)
				resp++
			}

		}
	}

	return resp
}

func dfs(grid [][]byte, row, col int) {

	// 如果超出范围则应当直接返回
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == '0' {
		return
	}
	//如果是1而且没有被标记过,那么说明是新陆地,将其标记为海洋

	grid[row][col] = '0'
	//如果是新发现的陆地那么应当继续尝试去探索周边陆地
	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)

	return

}
