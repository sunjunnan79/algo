package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)

	for i := range strs {

		runes := []rune(strs[i])

		for i := 0; i < len(runes)-1; i++ {
			for j := 0; j < len(runes)-1-i; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}

		str := string(runes)
		if _, ok := maps[str]; ok {
			maps[str] = append(maps[str], strs[i])
		} else {
			maps[str] = []string{strs[i]}
		}

	}

	var res [][]string
	for _, values := range maps {
		res = append(res, values)
	}
	return res
}
