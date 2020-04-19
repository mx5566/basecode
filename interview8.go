package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// [7,1,3,5,6,4]
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// 一直搜索如果遇到当前值小于下一个值的时候
	// 修改下一个值
	initvalue, lastvalue := prices[0], prices[0]
	var tempvale = 0
	var total = 0
	for _, value := range prices[1:] {
		tempvale = value

		if lastvalue < tempvale {
			lastvalue = tempvale
			continue
		}
		// 差价值
		total += lastvalue - initvalue
		initvalue = tempvale
		lastvalue = tempvale
	}

	// 原因是如果最后一个升序的序列不能处理到
	// 1,2,3,2,4,5
	if lastvalue > initvalue {
		total += lastvalue - initvalue
	}
	return total
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	minPrice := prices[0]
	array := make([]int, length)
	for i := 1; i < length; i++ {
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
		array[i] = int(math.Max(float64(array[i-1]), float64(prices[i]-minPrice)))
	}

	return array[length-1]
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	array := make([]int, length)
	array[0] = nums[0]
	maxValue := nums[0]
	for i := 1; i < length; i++ {
		if array[i-1] <= 0 {
			array[i] = nums[i]
		} else {
			array[i] = array[i-1] + nums[i]
		}

		if maxValue < array[i] {
			maxValue = array[i]
		}
	}

	return maxValue
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	array := make([]int, length)
	array[0] = nums[0]
	array[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < length; i++ {
		array[i] = int(math.Max(float64(array[i-1]), float64(nums[i]+array[i-2])))
	}

	return array[length-1]
}

type Solution struct {
	init    []int
	shuffle []int
}

func Constructor(nums []int) Solution {
	return Solution{
		init:    nums,
		shuffle: make([]int, 0),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.init
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.init)
	temp := make([]int, length)

	copy(temp, this.init)

	this.shuffle = this.shuffle[0:0]

	for length > 0 {
		n := rand.Intn(length)
		v := temp[n]

		temp = append(temp[0:n], temp[n+1:]...)

		this.shuffle = append(this.shuffle, v)

		length--
	}

	return this.shuffle
}

/////////////////////////////////////
type MinStack struct {
	stack     []int
	minStatck []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{stack: []int{},
		minStatck: []int{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.minStatck = append(this.minStatck, x)
		this.stack = append(this.stack, x)
	} else {
		currentMin := this.minStatck[len(this.minStatck)-1]
		if currentMin < x {
			this.minStatck = append(this.minStatck, currentMin)
		} else {
			this.minStatck = append(this.minStatck, x)
		}
	}
}

func (this *MinStack) Pop() {
	length := len(this.stack)
	if length <= 0 {
		return
	}

	this.stack = this.stack[0 : len(this.stack)-1]
	this.minStatck = this.minStatck[0 : len(this.minStatck)-1]
}

func (this *MinStack) Top() int {
	length := len(this.stack)
	if length <= 0 {
		return -1
	}

	return this.stack[length-1]
}

func (this *MinStack) GetMin() int {
	length := len(this.minStatck)
	if length <= 0 {
		return -1
	}

	return this.minStatck[len(this.minStatck)-1]
}

func fizzBuzz(n int) []string {
	if n == 0 {
		return []string{}
	}

	temp := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			temp = append(temp, "FizzBuzz")
		} else if i%3 == 0 {
			temp = append(temp, "Fizz")
		} else if i%5 == 0 {
			temp = append(temp, "Buzz")
		}

		temp = append(temp, strconv.Itoa(i))
	}

	return temp
}

func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}

	count := 0
	for i := 2; i <= n; i++ {

		if isSu(i) {
			count++
		}

	}

	return count
}

func isSu(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true

}

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) != 0 {
			count++
		}
	}

	return count
}

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor4() MyStack {
	return MyStack{queue: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	size := len(this.queue)

	if size == 0 {
		this.queue = append(this.queue, x)
		return
	}
	this.queue = append(this.queue, x)

	for size > 0 {
		value := this.queue[0]

		this.queue = this.queue[1:]

		this.queue = append(this.queue, value)
		size--
	}
}

//////////////////////

type MyQueue struct {
	a []int
	b []int
}

/** Initialize your data structure here. */
func Constructor5() MyQueue {
	return MyQueue{
		a: []int{},
		b: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	lengthb := len(this.b)
	if lengthb == 0 {
		lengtha := len(this.a)
		for lengtha > 0 {
			lengtha--
			value := this.a[lengtha]
			this.b = append(this.b, value)
		}

		// 清空a
		this.a = this.a[0:0]

	}

	lengthb = len(this.b)
	if lengthb == 0 {
		return 0
	}

	value := this.b[lengthb-1]
	this.b = this.b[0 : lengthb-1]

	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	lengthb := len(this.b)
	if lengthb == 0 {
		lengtha := len(this.a)
		for lengtha > 0 {
			lengtha--
			value := this.a[lengtha]
			this.b = append(this.b, value)
		}

		// 清空a
		this.a = this.a[0:0]

	}

	lengthb = len(this.b)
	if lengthb == 0 {
		return 0
	}

	return this.b[lengthb-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.a) == 0 && len(this.b) == 0
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.queue)
	if length == 0 {
		return 0
	}

	value := this.queue[0]

	this.queue = this.queue[1:]

	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	length := len(this.queue)
	if length == 0 {
		return 0
	}

	return this.queue[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	length := len(this.queue)
	if length == 0 {
		return true
	}

	return false
}

func isValid(s string) bool {
	for strings.Contains(s, "{}") || strings.Contains(s, "[]") || strings.Contains(s, "()") {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)

	}

	if s == "" {
		return true
	}

	return false
}

// 单调栈
// 单调递增 调调递减
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	length2 := len(nums2)
	stack := []int{}

	hash := map[int]int{}
	for i := 0; i < length2; i++ {
		lengthS := len(stack)
		for lengthS > 0 {
			lengthS--
			top := stack[lengthS]
			if top < nums2[i] {
				// 删除栈顶元素
				stack = stack[0:lengthS]

				// 记录比top大的第一个元素
				hash[top] = nums2[i]
			}
		}

		stack = append(stack, nums2[i])
	}

	lengthS := len(stack)
	for lengthS > 0 {
		lengthS--
		value := stack[lengthS]
		stack = stack[0:lengthS]
		hash[value] = -1
	}

	ret := []int{}
	length1 := len(nums1)
	for i := 0; i < length1; i++ {
		v, _ := hash[nums1[i]]
		ret = append(ret, v)
	}

	return ret
}

func calPoints(ops []string) int {
	total := 0
	stack := []int{}
	length := len(ops)
	for i := 0; i < length; i++ {
		value := ops[i]
		if value == "C" || value == "D" || value == "+" {
			lengthStack := len(stack)
			if lengthStack <= 0 {
				continue
			}
			if value == "C" {
				// 最后一个无效
				value := stack[lengthStack-1]
				stack = stack[0 : lengthStack-1]

				total -= value
			} else if value == "D" {
				temp := stack[lengthStack-1]
				stack = append(stack, temp*2)

				total += 2 * temp
			} else {
				if lengthStack == 1 {
					temp := stack[lengthStack-1]
					stack = append(stack, temp)

					total += temp
				} else if lengthStack >= 2 {
					temp1 := stack[lengthStack-1]
					temp2 := stack[lengthStack-2]
					stack = append(stack, temp1+temp2)

					total += temp1 + temp2
				}
			}
		} else {
			t, _ := strconv.Atoi(value)
			total += t
			stack = append(stack, t)
		}

	}

	return total
}

func backspaceCompare(S string, T string) bool {
	length1, length2 := len(S), len(T)

	s := []byte(S)
	t := []byte(T)

	for i := 0; i < length1; {
		if s[i] == '#' {
			if i > 0 {
				i -= 1

				s = append(s[0:i-1], s[i+1:]...)
				continue
			}

			s = append(s[0:i], s[i+1:]...)
		}
	}

	for i := 0; i < length2; {
		if t[i] == '#' {
			if i > 0 {
				i -= 1
				t = append(t[0:i-1], t[i+1:]...)
				continue
			}

			t = append(t[0:i], t[i+1:]...)
		}
	}

	S = string(s)
	T = string(t)

	return S == T
	//stack1 := []byte{}
	//stack2 := []byte{}
	//
	//length1, length2 := len(S),  len(T)
	//for i := 0; i < length1; i++ {
	//	lengthStack := len(stack1)
	//	if S[i] == '#' {
	//		if lengthStack > 0 {
	//			stack1 = stack1[0 : lengthStack-1]
	//		}
	//
	//		continue
	//	}
	//
	//	stack1 = append(stack1, S[i])
	//}
	//
	//for i := 0; i < length2; i++ {
	//	lengthStack := len(stack2)
	//	if T[i] == '#' {
	//		if lengthStack > 0 {
	//			stack2 = stack2[0 : lengthStack-1]
	//		}
	//
	//		continue
	//	}
	//
	//	stack2 = append(stack2, T[i])
	//}
	//
	//length1 = len(stack1)
	//length2 = len(stack2)
	//
	//if length1 != length2 {
	//	return false
	//}
	//
	//for i := 0; i < length1; i++ {
	//	if stack1[i] != stack2[i] {
	//		return false
	//	}
	//}
	//
	//return true
}

func RemoveOuterParentheses(S string) string {
	stack := []byte{}
	length := len(S)

	ret := []byte{}
	for i := 0; i < length; i++ {
		lengthS := len(stack)
		if S[i] == ')' {
			stack = stack[0 : lengthS-1]
			lengthS--
		}

		fmt.Println(ret)
		if lengthS > 0 {
			ret = append(ret, S[i])
		}

		if S[i] == '(' {
			stack = append(stack, S[i])
		}

	}

	fmt.Println(stack)
	return string(ret)
}

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length <= 0 || k <= 0 || k > length {
		return []int{}
	}

	// 队列存储索引
	queue := []int{}
	// 返回最大值列表
	ret := []int{}
	// 循环
	for i := 0; i < length; i++ {
		lengthQueue := len(queue)
		// 防止队列长度超过K的大小
		for lengthQueue > 0 && i-queue[0]+1 > k {
			queue = queue[1:lengthQueue]
			lengthQueue--
		}

		// 维持队列的单调性
		for lengthQueue > 0 {
			if nums[i] >= nums[queue[lengthQueue-1]] {
				queue = queue[0 : lengthQueue-1]
				lengthQueue--
			} else {
				break
			}
		}

		queue = append(queue, i)
		if i >= k-1 {
			ret = append(ret, nums[queue[0]])
		}

	}

	return ret
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}

	if k > len(arr) {
		return arr
	}

	sort.Ints(arr)

	return arr[0:k]
}
