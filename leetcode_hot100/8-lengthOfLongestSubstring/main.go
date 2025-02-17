package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("ggububgvfk"))
}

func lengthOfLongestSubstring(s string) int {
	lenth := 0
	start := 0
	runes := []rune(s)
	for start+lenth < len(runes) {
		c := confilct(runes[start : start+lenth+1])
		if c != -1 {

			start = start + c + 1
		} else {
			lenth++
		}
	}
	return lenth
}

func confilct(runes []rune) int {
	hash := make(map[rune]int)
	//去重
	for i := 0; i < len(runes); i++ {
		if v, ok := hash[runes[i]]; !ok {
			hash[runes[i]] = i
		} else {
			return v
		}
	}

	return -1
}
