package main

import (
	"fmt"
	"testing"
)

func TestInsertsort(t *testing.T) {
	arr := []int{5, 2, 4, 6, 1, 3}

	Insertsort(arr)
	fmt.Println(arr)
}
