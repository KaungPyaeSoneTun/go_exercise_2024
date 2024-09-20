package main

import "fmt"

func main() {
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

func removeDuplicates(nums []int) int {
	cacheMap := make(map[int]int)
	numValidEle := 0
	for i, num := range nums {
		if _, ok := cacheMap[num]; !ok {
			nums[numValidEle] = num
			numValidEle++
		}
		cacheMap[num] = i
	}
	fmt.Printf("Result Arr: %v \t Total Distinct Number: %v", nums, numValidEle)
	return numValidEle
}
