package main

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	//var q int
	//F(q) // 不能编译
}

func TestF1(t *testing.T) {
	var q []int
	F1(q) // 可以编译

	fmt.Printf("hello test %v\n", q)
}

func TestF2(t *testing.T) {
	var q []int
	F2(q) // 没问题
	//var q1 T1
	//F2(q1) // 无法编译 原因是q1是命名类型T1 而F2的参数要求是命名类型T2 命名类型 = 命名类型 不允许赋值操作
}
