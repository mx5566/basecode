package main

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {

	var nums = []int{1, 0, 2, 0, 0, 0, 5, 0, 0, 6, 0}
	MoveZeroes(nums)
	fmt.Println(nums)
}
