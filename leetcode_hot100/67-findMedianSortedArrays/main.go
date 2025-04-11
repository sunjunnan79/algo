package main

import "sort"

func main() {

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)
	sort.Ints(s)
	if len(s)%2 == 0 {
		return float64(s[len(s)/2-1]+s[len(s)/2]) / 2
	} else {
		return float64(s[len(s)/2])
	}
}
