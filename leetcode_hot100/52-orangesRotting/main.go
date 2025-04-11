package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}

func orangesRotting(grid [][]int) int {
	var rottedlist [][2]int
	var flag = true

	//一次遍历获取所有的烂橘子的位置
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rottedlist = append(rottedlist, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				flag = false
			}
		}
	}
	if flag {
		return 0
	}

	// 获取腐烂列表
	var getRottedList = func(rottedlist [][2]int) [][2]int {
		var l [][2]int
		for i := 0; i < len(rottedlist); i++ {
			var addToList = func(row, col int) {
				//如果超出范围或者已经腐烂,或者没有橘子则直接返回
				if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == 0 || grid[row][col] == 2 {
					return
				}
				grid[row][col] = 2
				//添加到最终边缘腐烂列表
				l = append(l, [2]int{row, col})
			}

			//尝试腐烂周围的所有橘子并得到腐烂橘子列表
			addToList(rottedlist[i][0]+1, rottedlist[i][1])
			addToList(rottedlist[i][0]-1, rottedlist[i][1])
			addToList(rottedlist[i][0], rottedlist[i][1]+1)
			addToList(rottedlist[i][0], rottedlist[i][1]-1)
		}
		return l
	}

	//尝试递归腐烂所有的橘子
	var count = -1
	for len(rottedlist) > 0 {
		rottedlist = getRottedList(rottedlist)
		count++
	}

	//遍历查找所有的橘子
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				//如果还有新鲜的橘子说明没有办法让所有的橘子腐烂
				return -1
			}
		}
	}

	return count
}
