// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  heap
 * @Version: 1.0.0
 * @Date: 2020/4/18 22:23
 */

package main

import (
	"fmt"
	"time"
)

// 定义类型
type ElementType int
type MaxHeap *HeapStruct

// 定义堆结构
type HeapStruct struct {
	//指向一个数组
	Element []ElementType
	//堆当前元素的个数
	size int
	//堆的最大容量
	capacity int
}

// 最大堆

func initHeap(arr []ElementType, size int, heap MaxHeap, maxCapacity int) {
	heap.capacity = maxCapacity

	heap.Element = make([]ElementType, maxCapacity+1)
	heap.size = size
	//第一个位置不存任何数据,交换结点时可以作为中间变量
	heap.Element[0] = ElementType(0)
	//构建完全二叉树
	for i := 0; i < size; i++ {
		heap.Element[i+1] = arr[i]
	}

	//寻找最后一个结点的父结点,作为初始值
	for pIndex := heap.size / 2; pIndex >= 1; pIndex-- {
		tmp := pIndex
		for (tmp << 1) <= heap.size { //表示该结点有孩子结点-->当该结点时叶子结点时,循环结束
			maxChildIndex := 0
			//寻找这个结点的最大子结点
			if (tmp<<1)+1 <= heap.size { //没有右孩子,则左孩子就是最大子结点
				maxChildIndex = (tmp << 1)
			} else { //从左右孩子中寻找最大子结点
				maxChildIndex = maxIndex(tmp<<1, (tmp<<1)+1, heap)
			}

			//比较最大子结点和当前父结点,如果父结点的值小于最大子结点的值,则交换两个结点
			if heap.Element[tmp] < heap.Element[maxChildIndex] {
				//交换两个结点
				heap.Element[0] = heap.Element[tmp]
				heap.Element[tmp] = heap.Element[maxChildIndex]
				heap.Element[maxChildIndex] = heap.Element[0]
				heap.Element[0] = 0
				tmp = maxChildIndex
			} else {
				break //当该结点不需要在交换时,结束向下查找
			}
		}
	}
}

func maxIndex(left, right int, heap MaxHeap) int {
	if heap.Element[left] > heap.Element[right] {
		return left
	} else {
		return right
	}
}

//原理:现在堆的最后增加一个结点,然后沿这堆树上升.
func maxHeapInsert(e ElementType, heap MaxHeap) int {
	//检查是否到达了堆的最大容量
	if heap.capacity == heap.size {
		return -1
	}

	heap.size++
	heap.Element[heap.size] = e
	for sIndex := heap.size; sIndex > 1; {
		//寻找这个结点的父结点
		pIndex := sIndex / 2
		if heap.Element[pIndex] < heap.Element[sIndex] {
			heap.Element[0] = heap.Element[pIndex]
			heap.Element[pIndex] = heap.Element[sIndex]
			heap.Element[sIndex] = heap.Element[0]
			heap.Element[0] = 0
			sIndex = pIndex
		} else {
			break
		}
	}

	return 0
}

//原理:将堆的最后的结点提到根结点，然后删除最大值，然后再把新的根结点向下进行调整,直到找到其符合的的位置.
func maxHeapPopE(heap MaxHeap, e *ElementType) int {
	fmt.Println(heap.size)
	if heap.size == 0 {
		return -1
	}
	*e = heap.Element[1]

	fmt.Println(*e)
	heap.Element[1] = heap.Element[heap.size]
	heap.size--
	pIndex := 1
	for pIndex<<1 <= heap.size { //有子结点
		maxChild := 0
		if (pIndex<<1)+1 > heap.size { //没有右孩子
			maxChild = pIndex << 1
		} else {
			maxChild = maxIndex(pIndex<<1, (pIndex<<1)+1, heap)
		}
		if heap.Element[pIndex] < heap.Element[maxChild] {
			heap.Element[0] = heap.Element[pIndex]
			heap.Element[pIndex] = heap.Element[maxChild]
			heap.Element[maxChild] = heap.Element[0]
			heap.Element[0] = 0
			pIndex = maxChild
		} else {
			break
		}
	}
	return 0
}

func main() {
	arr := []ElementType{5, 1, 13, 3, 16, 7, 10, 14, 6, 9}

	heap := new(HeapStruct)
	initHeap(arr /*sizeof(arr) / sizeof(int)*/, len(arr), heap, 20)

	fmt.Println(heap)

	maxHeapInsert(18, heap)

	var e ElementType
	maxHeapPopE(heap, &e)

	fmt.Println(fmt.Sprintf("e = %v", e))

	for i := 1; i <= heap.size; i++ {
		fmt.Println(heap.Element[i])
	}

	time.Sleep(time.Duration(10) * time.Second)
}
