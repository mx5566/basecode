package main

import (
	"fmt"
	"time"
)

// 输入: nums1 = [1,2,2,1], nums2 = [2,2]
// 输出: [2,2]
func Intersect(nums1 []int, nums2 []int) []int {
	var map1 map[int][]int
	var map2 map[int][]int

	map1 = map[int][]int{}
	map2 = map[int][]int{}

	for _, value := range nums1 {
		map1[value] = append(map1[value], value)
	}

	for _, value := range nums2 {
		map2[value] = append(map2[value], value)

	}

	var sli = []int{}
	for k, v := range map1 {
		value, ok := map2[k]
		if !ok {
			continue
		}

		sliTemp := []int{}
		if len(v) >= len(value) {
			sliTemp = value
		} else {
			sliTemp = v
		}

		sli = append(sli, sliTemp...)
	}

	return sli
}

func IntersectEx(nums1 []int, nums2 []int) []int {
	var map1 map[int]int

	map1 = map[int]int{}

	for _, value := range nums1 {
		map1[value]++
	}

	var sli = []int{}

	for _, value := range nums2 {
		if len(map1) <= 0 {
			break
		}

		count, ok := map1[value]

		if !ok {
			continue
		}

		if count <= 0 {
			delete(map1, value)
			continue
		}

		map1[value]--
		sli = append(sli, value)

	}

	return sli
}

type foo struct {
	Val int
}

type bar struct {
	Val int
}

func Equal() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}

	fmt.Print(a == b, c == foo(d), e == f)

	fmt.Printf("%T %v %p\n", a, a, a)
	fmt.Printf("%T %v %p\n", b, b, b)
	fmt.Printf("%T %v\n", foo(d), foo(d))
}

func A() int {
	fmt.Print("进入ASleep Before\n")
	time.Sleep(1000 * time.Millisecond)
	fmt.Print("进入ASleep After\n")

	return 1
}

func B() int {
	fmt.Print("进入BSleep Before\n")
	time.Sleep(200 * time.Millisecond)
	fmt.Print("进入BSleep After\n")

	return 2

}

func T11() {
	ch := make(chan int, 1)

	//ch <- A()

	go func() {
		select {
		case ch <- 4:
		case ch <- 6:

		case ch <- func() int {
			fmt.Print("进入ASleep Before\n")
			time.Sleep(1000 * time.Millisecond)
			fmt.Print("进入ASleep After\n")

			return 1
		}():
			fmt.Print("A")
		case ch <- func() int {
			fmt.Print("进入BSleep Before\n")
			time.Sleep(200 * time.Millisecond)
			fmt.Print("进入BSleep After\n")

			return 2

		}():
			fmt.Print("B")
		case ch <- 5:

			//default:
			//	ch <- 3
		}
	}()

	fmt.Println(<-ch)
}

func T11Get(t []byte) []byte {

	rr := t[0:3]
	raw := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1} //make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])

	rr[0] = 3
	return rr
}
