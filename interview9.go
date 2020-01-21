package main

import "fmt"

// 1s = 10 ^-9ns
// 0.1ms =
// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]
func Rotate(nums []int, k int) {

	var length = len(nums)
	if length <= 1 {
		return
	}

	if k >= length {

		k = k % length
	}

	if k <= 0 {
		return
	}

	fmt.Printf("----%d %d\n", k, length)
	copyNums := nums

	tempNums := copyNums[0 : length-k]
	copyNums = copyNums[length-k:]
	copyNums = append(copyNums, tempNums...)

	fmt.Printf("%v\n", nums)

	copy(nums, copyNums)

}
