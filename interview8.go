package main

// [7,1,3,5,6,4]
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// 一直搜索如果遇到当前值小于下一个值的时候
	// 修改下一个值
	initvalue, lastvalue := prices[0], prices[0]
	var tempvale = 0
	var total = 0
	for _, value := range prices[1:] {
		tempvale = value

		if lastvalue < tempvale {
			lastvalue = tempvale
			continue
		}
		// 差价值
		total += lastvalue - initvalue
		initvalue = tempvale
		lastvalue = tempvale
	}

	// 原因是如果最后一个升序的序列不能处理到
	// 1,2,3,2,4,5
	if lastvalue > initvalue {
		total += lastvalue - initvalue
	}
	return total
}
