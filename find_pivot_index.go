package main

import "fmt"

// https://leetcode.com/problems/find-pivot-index/?envType=study-plan&id=level-1
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	if nums[0] == nums[len(nums)-1] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[len(nums)-1]-nums[i] == nums[i-1] {
			return i
		}
	}

	return -1
}

func main() {
	//nums := []int{1, 7, 3, 6, 5, 6}
	nums := []int{2, 1, -1}
	fmt.Println(pivotIndex(nums))
}
