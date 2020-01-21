package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var arr = []int{1, 3, 3, 4, 4}
	len1 := RemoveDuplicates(arr)

	arr = arr[0:len1]

	fmt.Println(arr)
}
