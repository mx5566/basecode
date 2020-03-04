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
