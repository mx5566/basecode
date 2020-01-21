package main

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	fmt.Printf("--------%v\n", Intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func TestIntersectEx(t *testing.T) {
	fmt.Printf("--------Ex%v\n", IntersectEx([]int{1, 2, 2, 1}, []int{2, 2}))
}

func TestEqual(t *testing.T) {
	Equal()
}
func TestT11(t *testing.T) {
	T11()
}

func TestT11Get(t *testing.T) {
	var bb []byte = []byte{1, 1, 1, 1}

	var cc []byte

	fmt.Println(len(cc), cap(cc), cc)

	cc = bb // 值传递  吧bb的属性给了cc len  cap  ptr

	fmt.Println(len(cc), cap(cc), cc)

	fmt.Println(len(bb), cap(bb), &bb[0])

	data := T11Get(bb)
	fmt.Println(len(data), cap(data), &data[0], data)

	//fmt.Println("haha")
	//fmt.Printf("haha")
}
