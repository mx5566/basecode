package main

import (
	"fmt"
	"sort"
	"sync"
)

func ReverseString(s []byte) {
	var length = len(s)
	var l = length / 2

	for i := 0; i < l; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func ReverseInt(x int) int {
	y := 0
	for x != 0 {
		y = 10*y + x%10

		if y < -(1<<31) || y > ((1<<31)-1) {
			return 0
		}

		x /= 10
	}

	return y

	//
	//var min int32 = -1 * 0x80000000
	//var max int32 = 0x7fffffff
	//s := []byte{}
	//
	//var tempX = int(math.Abs(float64(x)))
	//
	//for tempX != 0 {
	//	k := byte(tempX % 10)
	//	s = append(s, k)
	//	tempX = tempX / 10
	//}
	//
	//var length = len(s)
	//var ret int
	//for i := 0; i < length; i++ {
	//	if s[i] == 0 {
	//		continue
	//	}
	//
	//	ret += int(s[i]) * int(math.Pow(float64(10), float64(length-i-1)))
	//}
	//
	//if x < 0 {
	//	ret *= -1
	//}
	//
	//fmt.Print(ret)
	//if ret < int(min) || ret > int(max) {
	//	return 0
	//}
	//
	//return ret
}

func FirstUniqChar(s string) int {
	var length = len(s)
	var arr = [256]int{}

	for i := 0; i < length; i++ {
		arr[s[i]]++
	}

	for i := 0; i < length; i++ {
		if arr[s[i]] == 1 {
			return i
		}
	}

	return -1
}

func IsAnagram(s string, t string) bool {
	var length1, length2 = len(s), len(t)

	if length1 != length2 {
		return false
	}

	var map1 = make(map[byte]int, length1)
	var map2 = make(map[byte]int, length1)

	for i := 0; i < length1; i++ {
		map1[s[i]]++
		map2[t[i]]++
	}

	for key, value1 := range map1 {
		var value2 int
		var ok bool
		if value2, ok = map2[key]; !ok {
			return false
		}

		if value1 != value2 {
			return false
		}
	}

	return true
}

func IsPalindrome(s string) bool {
	var length = len(s)

	var slice = []byte{}
	for i := 0; i < length; i++ {
		if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) {

			if s[i] >= 65 && s[i] <= 90 {
				slice = append(slice, s[i]+32)
				continue
			}

			slice = append(slice, s[i])
		}
	}

	length = len(slice)
	// A man, a plan, a canal: Panama
	for i := 0; i < length/2; i++ {
		if slice[i] != slice[length-i-1] {
			return false
		}
	}

	return true
}

func MyAtoi(str string) int {
	flg := 0
	res := 0
	startPos := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else {
			startPos = i
			if str[i] == '+' {
				flg = 1
				startPos += 1
			} else if str[i] == '-' {
				flg = -1
				startPos += 1
			}
			break
		}
	}

	fmt.Println(flg)
	for i := startPos; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
			if res > 1<<31-1 {
				break
			}
		} else {
			break
		}
	}

	fmt.Println(res)

	if flg < 0 {
		res = -res
	}
	if res > 1<<31-1 {
		return 1<<31 - 1
	}
	if res < -1<<31 {
		return -1 << 31
	}
	return res

	//// 去除空格
	//str = strings.TrimSpace(str)
	//
	//fmt.Println(str)
	//var length = len(str)
	//if length == 0 {
	//	return 0
	//}
	//
	//var neg = 1
	//var ch = str[0]
	//if ch != '+' && (ch < 48 || ch > 57) && ch != '-' {
	//	return 0
	//}
	//
	//if ch == '-' {
	//	neg = -1
	//}
	//
	//var isNum = false
	//if ch >= 48 && ch <= 57 {
	//	isNum = true
	//}
	//
	//var slice = []byte{}
	//if isNum {
	//	if str[0]-48 != 0 {
	//		slice = append(slice, str[0]-48)
	//	}
	//}
	//
	//var need = len(slice) <= 0
	//var con = false
	//for i := 1; i < length; i++ {
	//	if need {
	//		if i == 1 {
	//			if str[i] == 48 {
	//				con = true
	//				continue
	//			}
	//		}
	//
	//		if con == true && str[i] == 48 {
	//			continue
	//		}
	//
	//		con = false
	//
	//	}
	//
	//	if str[i] < 48 || str[i] > 57 {
	//		break
	//	}
	//
	//	slice = append(slice, str[i]-48)
	//}
	//
	//fmt.Println(slice)
	//
	//var ret = 0
	//length = len(slice)
	//
	//var min int32 = -1 * 0x80000000
	//var max int32 = 0x7fffffff
	//if length > 10 {
	//	if neg > 0 {
	//		return int(max)
	//	}
	//	return int(min)
	//}
	//
	//for _, ch := range slice {
	//	if ch > 9 {
	//		return 0
	//	}
	//	ret = ret*10 + int(ch)
	//}
	//
	//fmt.Println(ret)
	//
	//ret = ret * neg
	//if ret < int(min) {
	//	return int(min)
	//}
	//
	//if ret > int(max) {
	//	return int(max)
	//}
	//
	//return ret
}

//hello ll
func StrStr(haystack string, needle string) int {
	needleLength := len(needle)
	length := len(haystack)
	for i := 0; i <= length-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}
	return -1

	//var length1 = len(haystack)
	//var length2 = len(needle)
	//
	//if length2 <= 0 {
	//	return 0
	//}
	//
	//if length1 < length2 {
	//	return -1
	//}
	//
	//var src string
	//for i := 0; i < length1; i++ {
	//	src = haystack[i:]
	//
	//	leng := len(src)
	//
	//	for j := 0; j < length2 && j < leng; j++ {
	//		if needle[j] != src[j] {
	//			break
	//		}
	//
	//		if j+1 == length2 {
	//			return i
	//		}
	//	}
	//}
	//
	//return -1
}

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := "1"
	for i := 1; i < n; i++ {
		var last = str[0]
		var nTemp = 1
		str1 := ""
		length := len(str)
		for j := 1; j < length; j++ {
			if last == str[j] {
				nTemp++
				continue
			}

			//str1 += strconv.Itoa(nTemp)
			//str1 += strconv.Itoa(int(last - 48))
			str1 += fmt.Sprintf("%d%c", nTemp, last)

			last = str[j]
			nTemp = 1
		}

		str1 += fmt.Sprintf("%d%c", nTemp, last)

		//str1 += strconv.Itoa(nTemp)
		//str1 += strconv.Itoa(int(last - 48))

		str = str1
	}

	return str
}

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	if length == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	var first = strs[0]
	var last = strs[len(strs)-1]

	var indexRet = -1

	var length1 = len(first)
	var length2 = len(last)
	for i := 0; i < length1 && i < length2; i++ {
		if first[i] != last[i] {
			break
		}

		indexRet = i
	}

	if indexRet == -1 {
		return ""
	}

	if indexRet == 0 {
		return string(first[indexRet])
	}

	return first[0 : indexRet+1]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 我觉得这个题目没多大意义
// 和我们通常所说的链表不是一个概念，违背链表的逻辑le
func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 空间复杂度o(n)
// 增加一个o(n)的空间
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	next := head
	slice := []*ListNode{}
	for next != nil {
		slice = append(slice, next)
		next = next.Next
	}

	var length = len(slice)
	if length == 0 {
		return nil
	}

	if n == length {
		return head.Next
	}

	preNode := slice[length-n-1]
	curNode := slice[length-n]

	preNode.Next = curNode.Next
	curNode = nil

	return slice[0]
}

func ReverseList(head *ListNode) *ListNode {
	// 遞歸算法
	if head == nil || head.Next == nil {
		return head
	}

	node := ReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return node

	// 1原地算法
	// 用两个额外的节点 一个用来返回
	// 一个用来作为中间节点
	//var newHead *ListNode = nil
	//var node *ListNode = nil
	//
	//for head != nil {
	//	node = head
	//	head = head.Next
	//
	//	node.Next = newHead
	//	newHead = node
	//}
	//
	//return newHead

	// 2
	//next := head
	//slice := []*ListNode{}
	//for next != nil {
	//	slice = append(slice, next)
	//	next = next.Next
	//}
	//
	//var length = len(slice)
	//if length == 0 {
	//	return nil
	//}
	//
	//for i := 0; i < length/2; i++ {
	//	slice[i], slice[length-1-i] = slice[length-1-i], slice[i]
	//}
	//
	//next = slice[0]
	//for i := 1; i < length; i++ {
	//	next.Next = slice[i]
	//	next = next.Next
	//}
	//
	//next.Next = nil
	//
	//return slice[0]
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var lT = &ListNode{}
	var lRet = lT

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			lTemp := new(ListNode)
			lTemp.Val = l1.Val
			lTemp.Next = nil

			lT.Next = lTemp
			lT = lT.Next

			l1 = l1.Next
		} else {
			lTemp := new(ListNode)
			lTemp.Val = l2.Val
			lTemp.Next = nil

			lT.Next = lTemp
			lT = lT.Next

			l2 = l2.Next
		}
	}

	for l1 != nil {
		lTemp := new(ListNode)
		lTemp.Val = l1.Val
		lTemp.Next = nil

		lT.Next = lTemp
		lT = lT.Next

		l1 = l1.Next
	}

	for l2 != nil {
		lTemp := new(ListNode)
		lTemp.Val = l2.Val
		lTemp.Next = nil

		lT.Next = lTemp
		lT = lT.Next

		l2 = l2.Next
	}

	return lRet.Next
}

func IsPalindromeNode(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 查找中间节点
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	initHead := head
	newHead := ReverseList(slow.Next)
	// 遍历一遍
	for initHead != nil {

		if initHead.Val != newHead.Val {
			return false
		}

		initHead = initHead.Next
		newHead = newHead.Next
	}

	return true
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		if fast == slow {
			return true
		}

		slow = slow.Next
	}

	return false
}

func Sum(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	if len(arr) == 1 {
		return arr[0], 1
	}

	sumRet := arr[0]
	count := 1

	sum1, count1 := Sum(arr[1:])

	sumRet += sum1
	count += count1
	return sumRet, count
}

func FindMaxNum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	maxNum := arr[0]

	ret := FindMaxNum(arr[1:])
	if maxNum < ret {
		maxNum = ret
	}

	return maxNum
}

func BinarySerach(arr []int, low, high, item int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}

	if low > high {
		return -1
	}

	ret := -1
	index := (low + high) / 2

	fmt.Println(low, " ", high, " ", index)

	if item == arr[index] {
		return index
	} else if item > arr[index] {
		ret = BinarySerach(arr, index+1, high, item)
	} else {
		ret = BinarySerach(arr, low, index-1, item)
	}

	return ret
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	baseValue := arr[0]
	leftArr := []int{}
	rightArr := []int{}

	for _, value := range arr[1:] {
		if value > baseValue {
			rightArr = append(rightArr, value)
			continue
		}

		leftArr = append(leftArr, value)
	}

	left := QuickSort(leftArr)

	right := QuickSort(rightArr)

	retArr := make([]int, 0)
	retArr = append(left, baseValue)
	retArr = append(retArr, right...)

	return retArr
}

// 这些是伪代码
type Account struct {
	sync.Mutex
	nMoney     uint64
	nAccountID uint64
}

// one to one
func TranToHashCode(account Account) int {
	return 0
}

// 取款
func WithDraw(account Account, nAmount uint64) {

}

// 存款
func Deposit(account Account, nAmount uint64) {

}

// 产生死锁的方式
// 如果有两个携程调用的话，并且调用顺序为
// Transaction(from, to,10)
// Transaction(to, from,10)
// 由于并发导致的交叉执行会产生死锁的可能
func Transaction(from Account, to Account, aMount uint64) {
	from.Lock()
	to.Lock()

	WithDraw(from, aMount)
	Deposit(to, aMount)

	to.Unlock()
	from.Unlock()
}

//

// 解决死锁的方式
// 我们 通过给每个账号求一个hash值人为的改变顺序
func Transaction2(from Account, to Account, aMount uint64) {
	var fromHashCode int = TranToHashCode(from)
	var toHashCode int = TranToHashCode(to)

	if fromHashCode < toHashCode {
		from.Lock()
		to.Lock()

		WithDraw(from, aMount)
		Deposit(to, aMount)

		to.Unlock()
		from.Unlock()
		return
	}

	to.Lock()
	from.Lock()

	WithDraw(from, aMount)
	Deposit(to, aMount)

	from.Unlock()
	to.Unlock()
}
