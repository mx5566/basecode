package main

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	fmt.Print(ContainsDuplicate([]int{1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8}))
}

func TestSingleNumber(t *testing.T) {
	fmt.Print(SingleNumber([]int{1, 1, 1, 2, 2, 3}))
}
