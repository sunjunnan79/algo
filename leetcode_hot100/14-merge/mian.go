package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 1}, {0, 0}}))
}

// 又被gpt爆了
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 按区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// 如果当前区间和下一个区间有重叠，合并它们
		if current[1] >= intervals[i][0] {
			current[1] = max(current[1], intervals[i][1])
		} else {
			// 否则，当前区间不能合并，保存并更新当前区间
			result = append(result, current)
			current = intervals[i]
		}
	}

	// 最后将剩余的区间加入结果
	result = append(result, current)
	return result
}
