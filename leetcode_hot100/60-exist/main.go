package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}},
		"ADD"))
}

func exist(board [][]byte, word string) bool {

	var feedback func(board [][]byte, used [][]bool, word string, i [2]int, j int) bool
	var used = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[i]))
	}

	feedback = func(board [][]byte, used [][]bool, word string, i [2]int, j int) bool {

		if i[0] < 0 || i[0] >= len(board) || //超范围
			i[1] < 0 || i[1] >= len(board[i[0]]) || //超范围
			used[i[0]][i[1]] || // 已经使用
			board[i[0]][i[1]] != byte(rune(word[j])) { //比较是否相同
			//如果已经使用过直接返回false
			return false
		}

		if j+1 == len(word) {
			return true
		}
		//如果为使用过则标记
		used[i[0]][i[1]] = true

		//尝试进行查找,如果其中有一个为ture则直接返回true
		if feedback(board, used, word, [2]int{i[0] - 1, i[1]}, j+1) ||
			feedback(board, used, word, [2]int{i[0] + 1, i[1]}, j+1) ||
			feedback(board, used, word, [2]int{i[0], i[1] + 1}, j+1) ||
			feedback(board, used, word, [2]int{i[0], i[1] - 1}, j+1) {
			return true
		}

		//没有则回溯
		used[i[0]][i[1]] = false
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if feedback(board, used, word, [2]int{i, j}, 0) {
				return true
			}
		}
	}
	return false
}
