package main

func main() {

}

func majorityElement(nums []int) int {
	var hashmap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
	}
	for k, v := range hashmap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
