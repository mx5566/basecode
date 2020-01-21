package main

import (
	"fmt"
	"testing"
)

func TestMartixRotate(t *testing.T) {
	var arr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	MartixRotate(arr)
	fmt.Println(arr)
}
