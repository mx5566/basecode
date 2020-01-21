package main

//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
func MoveZeroes(nums []int) {
	var nZero = 0
	for index, value := range nums {
		if value == 0 {
			nZero++
			continue
		}

		if nZero == 0 {
			continue
		}

		nums[index-nZero] = value
		nums[index] = 0

	}
}
