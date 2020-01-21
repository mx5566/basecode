package main

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	value := MaxProfit([]int{1, 2, 3, 4, 5})

	fmt.Printf("maxProfit : %d\n", value)
}

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := MaxProfit([]int{1, 2, 3, 4, 5})

		fmt.Printf("maxProfit : %d\n", value)
	}

}
