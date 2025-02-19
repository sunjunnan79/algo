package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("adsrtabb"))
}

func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	start, end := 0, 0

	for i, c := range s {
		//记录当前最远的
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}

		// 如果遇到重复值则添加到结果中去
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
