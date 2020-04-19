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

func TestConstructor(t *testing.T) {
	obj := Constructor([]int{})
	param_1 := obj.Reset()
	param_2 := obj.Shuffle()

	fmt.Println(param_1, " ", param_2)
}

func TestConstructor1(t *testing.T) {
	obj := Constructor1()
	obj.Pop()
	fmt.Println(obj.Top())

	//fmt.Println(param_1, " ", param_2)
}

func TestRemoveOuterParentheses(t *testing.T) {
	fmt.Println(RemoveOuterParentheses("(())"))
}
