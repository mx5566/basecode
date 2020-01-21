package main

func ContainsDuplicate(nums []int) bool {

	var maps map[int]int

	maps = map[int]int{}

	for _, value := range nums {
		maps[value]++

		if maps[value] >= 2 {
			return true
		}
	}
	return false
}

func SingleNumber(nums []int) int {
	a := 0
	for _, i := range nums {
		a ^= i
	}
	return a

	//
	//sort.Ints(nums)
	//
	//var length = len(nums)
	//
	//if length == 1 {
	//	return nums[0]
	//}
	//var count = 1
	//var ret = 0
	//
	//for i := 1; i < length; i++ {
	//	if nums[i] == nums[i-1] {
	//		count++
	//		continue
	//	}
	//
	//	if count == 1 {
	//		ret = nums[i-1]
	//		count = 0
	//		break
	//	}
	//
	//	count = 1
	//
	//}
	//
	//if count == 1 {
	//	ret = nums[length-1]
	//}
	//return ret
}
