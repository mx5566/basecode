package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6}

	Rotate(nums, 3)

	fmt.Printf("%v\n", nums)
}
