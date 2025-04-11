package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("cars", []string{"car", "ca", "rs"}))
}

func wordBreak(s string, wordDict []string) bool {
	var hash = make(map[uint8][]string)
	var flaghash = make(map[int]bool)
	//获取倒排表
	for i := 0; i < len(wordDict); i++ {
		if len(wordDict[i]) > 0 {
			hash[wordDict[i][0]] = append(hash[wordDict[i][0]], wordDict[i])
		}
	}

	var feedback func(str string) bool
	feedback = func(str string) bool {
		if flaghash[len(str)] {
			return false
		}

		if str == s {
			return true
		}
		for _, v := range hash[s[len(str)]] {
			if strings.HasPrefix(s, str+v) {
				if feedback(str + v) {
					return true
				} else {
					flaghash[len(str+v)] = true
				}
			}
		}
		return false
	}

	return feedback("")
}

//func wordBreak(s string, wordDict []string) bool {
//	var feedback func(temp string) bool
//	feedback = func(temp string) bool {
//		if temp == s {
//			return true
//		}
//
//		for i := 0; i < len(wordDict); i++ {
//			if strings.HasPrefix(s, temp+wordDict[i]) {
//				if feedback(temp+wordDict[i]) == true {
//					return true
//				}
//			}
//
//		}
//
//		return false
//	}
//
//	return feedback("")
//}
