package main

import (
	"fmt"
)

// https://leetcode.com/problems/running-sum-of-1d-array/?envType=study-plan&id=level-1
func runningSum(nums []int) []int {
	var sumArray []int
	previousValue := nums[0]
	sumArray = append(sumArray, nums[0])

	for i := 1; i < len(nums); i++ {
		previousValue = nums[i] + previousValue
		sumArray = append(sumArray, previousValue)
	}

	return sumArray
}

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(runningSum(numbers))
}
