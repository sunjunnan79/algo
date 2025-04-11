package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("12345", "54321"))

}

func longestCommonSubsequence(text1 string, text2 string) int {
	var dp = make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

//回溯法会直接超时
//func longestCommonSubsequence(text1 string, text2 string) int {
//	var m int
//	hashmap1 := make(map[string]bool)
//	hashmap2 := make(map[string]bool)
//	feedback := func(temp string, s string, m map[string]bool) {}
//	feedback = func(temp string, s string, m map[string]bool) {
//		for i := 0; i < len(s); i++ {
//			m[temp+s[i:i+1]] = true
//			feedback(temp+s[i:i+1], s[i+1:], m)
//		}
//	}
//	feedback("", text1, hashmap1)
//	feedback("", text2, hashmap2)
//	for k := range hashmap1 {
//		if hashmap2[k] {
//			m = max(m, len(k))
//		}
//	}
//	return m
//}
