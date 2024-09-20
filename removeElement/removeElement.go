package main

import "fmt"

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
}

// return the valid numbers of ele in the array not the array size itself
// replace the target element  with valid element
func removeElement(nums []int, target int) int {
	validEleSize := 0
	for i := range nums {
		if nums[i] != target {
			nums[validEleSize] = nums[i]
			validEleSize++
		}
	}
	fmt.Printf("The result array is %v \n", nums)
	return validEleSize
}
