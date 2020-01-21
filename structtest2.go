package main

import "fmt"

var slice1 []int
var name string

type struct1 struct {
	name    string  // 普通变量
	namePtr *string // 指针变量
	sli     []int   // slice
	sliPtr  *[]int  // slice指针
}

func main() {
	name = "mPtr"
	slice1 = []int{4, 5, 6}
	var s1 struct1
	s1.name = "mx"
	s1.namePtr = &name
	s1.sli = []int{1, 2, 3}
	s1.sliPtr = &slice1

	var s2 *struct1
	s2 = &s1

	fmt.Println("s1--- %s %p %v %v", s1.name, s1.namePtr, s1.sli, s1.sliPtr)
	fmt.Println("s2--- %s %p %v %v", s2.name, s2.namePtr, s2.sli, s2.sliPtr)

	var name1 string
	var slice2 []int
	slice2 = []int{10, 11, 12}
	name1 = "s2mPtr"
	s2.name = "s2mx"
	s2.namePtr = &name1
	s2.sli = []int{7, 8, 9}
	//s2.sli[0] = 9
	s2.sliPtr = &slice2

	fmt.Println("after s1--- %s %p %v %v", s1.name, s1.namePtr, s1.sli, s1.sliPtr)
	fmt.Println("after s2--- %s %p %v %v", s2.name, s2.namePtr, s2.sli, s2.sliPtr)

	var slice3 []int
	slice3 = slice2

	fmt.Println("before slice3--- %v", slice3)

	slice3 = slice1

	fmt.Println("after slice3--- %v %v", slice2, slice3)

	var arr = [3]int{1, 2, 3}
	var arr1 = [3]int{2, 4, 6}
	slice3 = arr[0:3] // 一个浅拷贝
	fmt.Println("after after slice3--- %v arr %v", slice3, arr)

	//slice3[0] = 10
	//arr[0] = 11
	//var slice4 *[]int

	slice3 = arr1[0:3]
	fmt.Println("after after slice3--- %v arr %v", slice3, arr)

	//slice4 = &slice3
	testTran(slice3)
	fmt.Println("after after slice3--- %v", slice3)

}

func testTran(a []int) {
	fmt.Println("after after slice3--- %v", a)

	a = []int{1, 3, 5}[0:3]

	fmt.Println("after after slice3--- %v", a)

}
