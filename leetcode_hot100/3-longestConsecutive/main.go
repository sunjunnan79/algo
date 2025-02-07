package main

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
}

func longestConsecutive(nums []int) int {

	var hash = make(map[int]int)

	// 建立倒排表
	for i := range nums {
		hash[nums[i]] = 0
	}

	maxNum := 0

	for key := range hash {
		//当且仅当 前一个数不存在的时候才进行遍历
		if _, ok := hash[key-1]; !ok {
			//计算当前的数量,起始数量显然为1,因为当前的是确定存在的
			currentkey := key
			currentNum := 1

			for {
				if _, ok := hash[currentkey+1]; !ok {
					break
				} else {
					currentkey++
					currentNum++
				}
			}

			if currentNum > maxNum {
				maxNum = currentNum
			}
		}
	}

	return maxNum
}
