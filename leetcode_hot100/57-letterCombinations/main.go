package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var hash = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var traceback func(digits []rune, path []rune, used []bool)
	var res []string
	var l = len(digits)
	var used = make([]bool, len(digits))

	traceback = func(digits []rune, path []rune, used []bool) {

		if len(path) == l {
			if len(path) == 0 {
				return
			}

			res = append(res, string(path))
			return
		}

		for _, s := range hash[digits[len(path)]] {
			traceback(digits, append(path, s), used)
		}

	}

	traceback([]rune(digits), []rune{}, used)

	return res
}
