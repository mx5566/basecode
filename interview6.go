package main

func RemoveDuplicates(nums []int) int {
	//	var m runtime.MemStats

	len1 := len(nums)

	if len1 <= 1 {
		return len1
	}

	var temp = nums[0]
	for i := 1; i < len1; {
		var temp2 = nums[i]

		// 下一个元素和上一个元素不相等
		// 修改上一个元素
		if temp2 != temp {
			i++
			temp = temp2

			continue
		}

		if i+1 == len1 {
			nums = nums[:i]
			break
		}

		// 如果和上一个元素相等 删除这个元素i不递增
		nums = append(nums[:i], nums[i+1:len1]...)
		len1 = len(nums)

	}

	//	runtime.ReadMemStats(&m)
	//	gb := 1024 * 1024 * 1024.0
	//	logstr := fmt.Sprintf("\nAlloc = %v\tTotalAlloc = %v\tSys = %v\t NumGC = %v\n\n", float64(m.Alloc)/gb, float64(m.TotalAlloc)/gb, float64(m.Sys)/gb, m.NumGC)

	//	log.Print(logstr)
	return len(nums)
}
