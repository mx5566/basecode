package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	var arr = []byte{'h', 'e', 'l', 'l', 'o'}

	ReverseString(arr)

	fmt.Printf("%v", arr)
}

func TestReverseInt(t *testing.T) {
	fmt.Print(ReverseInt(123))
}

func TestIsPalindrome(t *testing.T) {
	fmt.Print(IsPalindrome("A man, a plan, a canal: Panama"))
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(MyAtoi("42"))
}

func TestStrStr(t *testing.T) {
	fmt.Println(StrStr("lello", ""))
}

func TestCountAndSay(t *testing.T) {
	fmt.Println(CountAndSay(5))
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	//fmt.Println(LongestCommonPrefix([]string{"c", "c"}))

}

func TestRemoveNthFromEnd(t *testing.T) {
}

func TestSum(t *testing.T) {
	fmt.Println(Sum([]int{1, 1, 2}))
}

func TestFindMaxNum(t *testing.T) {
	fmt.Println(FindMaxNum([]int{1, 6, 3}))
}

func TestBinarySerach(t *testing.T) {
	fmt.Println(BinarySerach([]int{1, 4, 5}, 0, 2, 0))
	//fmt.Println(BinarySerach([]int{1, 4, 5}, 0, 2, 0))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{1, 5, 2, 3}))
}
