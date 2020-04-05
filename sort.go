package main

func Insertsort(arr []int) {
	length := len(arr)

	for i := 1; i < length; i++ {
		key := arr[i]

		j := i - 1
		for j >= 0 && arr[j] <= key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// 把剩余的元素合并到
	copy(nums1[0:p2+1], nums2[0:p2+1])
}

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	if n <= 0 {
		return -1
	}

	if n == 1 {
		return 1
	}

	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
			continue
		}

		left = mid + 1
	}

	return left
}
