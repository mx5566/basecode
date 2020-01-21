package main

// [1,2,3] 123 + 8  = 131
// [9,9,9] [1,0,0,0]
func PlusOne(digits []int) []int {
	var length = len(digits)

	for i := length - 1; i >= 0; i-- {
		digits[i]++

		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	sli := make([]int, length+1)
	sli[0] = 1

	return sli
}
