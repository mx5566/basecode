package main

//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func TwoSum(nums []int, target int) []int {

	var length = len(nums)
	//for i := 0; i < length; i++ {
	//	for j := i + 1; j < length; j++ {
	//		if nums[i]+nums[j] != target {
	//			continue
	//		}
	//
	//		return []int{i, j}
	//	}
	//}
	//
	//return nil

	maps := make(map[int]int)
	for k := 0; k < length; k++ {
		if index, ok := maps[target-nums[k]]; ok {
			return []int{index, k}
		}

		maps[nums[k]] = k
	}

	return nil
}
