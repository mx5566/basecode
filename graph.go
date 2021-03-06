package main

import (
	"fmt"
	"math"
	"strconv"
)

// 逻辑不是很严谨  越界的没考虑-- scanf


// 边节点结构
type EdgeTableNode struct {
	index         int            // 顶点索引
	weight        int            // 权重
	edgeTableNode *EdgeTableNode // 下一个临界点的指针
}

// 顶点的数据信息
type VertInfo struct {
	value int
	name  string
}

// 顶点节点
type VertNode struct {
	vertInfo      VertInfo //  定点的数据信息
	edgeTableNode *EdgeTableNode
}

type Graph struct {
	vertNum  int
	grapType uint8
	edgeNum  int

	hasCircle bool
	allCircle [][]int
	visted    []int

	vertNode []*VertNode
}

var arrayList []int

func NewGraph(vertNum int, edgeNum int, grapType uint8) *Graph {
	return &Graph{vertNum: vertNum, hasCircle: false, allCircle: [][]int{},
		visted: make([]int, vertNum), grapType: grapType,
		edgeNum: edgeNum, vertNode: []*VertNode{}}
}

// 只做了有向图的初始化
// 没考虑无向图
func (this *Graph) InitGrap() {
	// 顶点初始化
	for i := 0; i < this.vertNum; i++ {
		vert := &VertNode{}

		vert.vertInfo.value = i
		vert.vertInfo.name = "V" + strconv.Itoa(i)

		fmt.Println(*vert)

		this.vertNode = append(this.vertNode, vert)
	}

	// 边初始化
	var startVert int
	var endVert int
	var weight int
	var n int
	for i := 0; i < this.edgeNum; i++ {
		n, _ = fmt.Scanf("%d %d %d", &startVert, &endVert, &weight)

		fmt.Printf("%d %d %d\n", startVert, endVert, n)

		var edgeNode = this.vertNode[startVert].edgeTableNode
		if edgeNode == nil {
			tempNode := new(EdgeTableNode)
			tempNode.weight = weight
			tempNode.index = endVert
			tempNode.edgeTableNode = nil
			this.vertNode[startVert].edgeTableNode = tempNode
			continue
		}

		for edgeNode != nil {
			// 单链表尾插节点
			if edgeNode.edgeTableNode == nil {
				tempNode := new(EdgeTableNode)
				tempNode.weight = weight
				tempNode.index = endVert
				tempNode.edgeTableNode = nil
				edgeNode.edgeTableNode = tempNode
				break
			}

			edgeNode = edgeNode.edgeTableNode
		}
	}
}

func FindArray(arr []int, v int) int {
	for index, value := range arr {
		if v == value {
			return index
		}
	}

	return -1
}

func (this *Graph) PrintAllCircle() {
	fmt.Println("打印环")
	for _, sli := range this.allCircle {
		fmt.Println(sli)
	}
}

// 遍历整个图使用深度优先算法
func (this *Graph) Dfs(v int) {
	if this.visted[v] == 1 {
		return
	}

	this.visted[v] = 1

	edge := this.vertNode[v].edgeTableNode
	for edge != nil {
		this.Dfs(edge.index)
		edge = edge.edgeTableNode
	}
}

// 遍历图
func (this *Graph) DfsAll() {
	for i := 0; i < this.vertNum; i++ {
		this.Dfs(i)
	}
}

// 通过迭代的方式遍历 --栈结构
//
func (this *Graph) DfsStatck() {
	var stack = CreateStack()

	for i := 0; i < this.vertNum; i++ {
		if this.visted[i] == 1 {
			continue
		}

		stack.Push(i)
		this.visted[i] = 1
		for !stack.IsEmpty() {
			elemet := stack.Get()

			node := this.vertNode[elemet.(int)].edgeTableNode

			for node != nil && this.visted[node.index] == 1 {
				node = node.edgeTableNode
			}

			if node != nil && this.visted[node.index] == 0 {
				this.visted[node.index] = 1
				stack.Push(node.index)
			} else {
				stack.Pop()
			}
		}
	}

}

// 查找所有换
func (this *Graph) FindCircle() {
	// 初始化0 表示没有被查找过
	this.DfsCircle(0)
}

// 遍历查找环
func (this *Graph) DfsCircle(v int) {
	j := FindArray(arrayList, v)
	if j != -1 {
		this.hasCircle = true

		tempSlice := make([]int, len(arrayList)-j)
		copy(tempSlice, arrayList[j:])
		this.allCircle = append(this.allCircle, tempSlice)
		return
	}

	arrayList = append(arrayList, v)

	edge := this.vertNode[v].edgeTableNode
	for edge != nil {
		this.DfsCircle(edge.index)
		edge = edge.edgeTableNode
	}

	// 如果走到这里说明当前这个节点不是环路点
	// 移除掉
	arrayList = arrayList[0 : len(arrayList)-1]
}

func GrapBfs() {
	var grap = NewGraph(5, 6, 1)
	grap.InitGrap()

	grap.BfsNearPath(0, 4)
}

func GrapTuoPuSort() {
	var grap = NewGraph(5, 6, 1)
	grap.InitGrap()

	grap.TuoPuSort()
}

func GrapTest() {
	var grap = NewGraph(4, 5, 1)
	grap.InitGrap()

	grap.FindCircle()
	grap.PrintAllCircle()

}

func GrapDfs() {
	var grap = NewGraph(4, 5, 1)
	grap.InitGrap()

	grap.DfsAll()
}

// 利用栈进行遍历图
func GrapStack() {

	var grap = NewGraph(4, 5, 1)
	grap.InitGrap()

	grap.DfsStatck()
}

func GrapDijstra() {
	var grap = NewGraph(4, 5, 1)
	grap.InitGrap()

	grap.Dijkstra(0, 3)
}

func main() {
	//GrapDfs()
	//GrapTuoPuSort()
	//GrapBfs()
	GrapDijstra()
}

// 队列
/*
队列的特性较为单一，基本操作即初始化、获取大小、添加元素、移除元素等。
最重要的特性就是满足先进先出
*/
type linkNode struct {
	value MapParent
	next  *linkNode
}

type linkedList struct {
	head  *linkNode
	tail  *linkNode
	count int
}

func NewLinkList() *linkedList {
	return &linkedList{head: nil, tail: nil, count: 0}
}

func (this *linkedList) IsEmpty() bool {
	return this.count == 0
}

func (this *linkedList) Add(value MapParent) {
	node := new(linkNode)
	node.value = value

	this.count++
	if this.tail == nil {
		this.head = node
		this.tail = node
		node.next = nil
		return
	}

	this.tail.next = node
	node.next = nil
	this.tail = node
}

func (this *linkedList) Delete() *linkNode {
	if this.head == nil {
		return nil
	}

	this.count--
	if this.head == this.tail {
		node := this.head
		this.head = nil
		this.tail = nil

		return node
	}

	node := this.head
	this.head = this.head.next

	return node
}

type Queue struct {
	link *linkedList
}

func NewQueue() *Queue {
	return &Queue{link: NewLinkList()}
}

//加入队列
func (this *Queue) Put(value MapParent) {
	this.link.Add(value)
}

//pop出队列
func (this *Queue) Pop() *linkNode {
	return this.link.Delete()
}

//获得队列的长度
func (this *Queue) GetSize() int {
	return this.link.count
}

func (this *Queue) IsEmpty() bool {
	return this.GetSize() == 0
}

// 初始化队列
var queue *Queue = NewQueue()

type MapParent struct {
	parent int
	son    int
}

// 查找最短路径
func (this *Graph) BfsNearPath(start, end int) []int {
	if start == end {
		return []int{start}
	}

	// 用来存储的是找到终点之前所有出队列的元素
	mapParent := make([]MapParent, 0)

	// 根据初始节点 把他的邻接点放入队列
	node := this.vertNode[start]
	for node.edgeTableNode != nil {
		index := node.edgeTableNode.index

		queue.Put(MapParent{parent: start, son: index})
		node.edgeTableNode = node.edgeTableNode.edgeTableNode
	}

	var find = false
	for !queue.IsEmpty() {
		// 检测队列的元素
		node := queue.Pop()

		// 已经被查看过得元素不需要再次查询 防止出现死循环
		if this.visted[node.value.son] == 1 {
			continue
		}

		// 记录出队的元素
		mapParent = append(mapParent, node.value)

		if node.value.son == end {
			find = true
			break
		}

		// 节点的邻接点入队列
		grapNode := this.vertNode[node.value.son]
		for grapNode.edgeTableNode != nil {
			index := grapNode.edgeTableNode.index
			// 记录父节点与子节点的映射
			queue.Put(MapParent{parent: node.value.son, son: index})
			grapNode.edgeTableNode = grapNode.edgeTableNode.edgeTableNode
		}

		// 记录被查询过
		this.visted = append(this.visted, node.value.son)
	}

	// 逆序查找路径
	path := []int{}

	if find == true {
		path = append(path, end)
		son := end
		for i := len(mapParent) - 1; i >= 0; i-- {
			if son == mapParent[i].son {
				path = append(path, mapParent[i].parent)
				son = mapParent[i].parent
			}
		}
	}

	// 打印查找的路径
	fmt.Println(path)
	return path
}

// 拓扑排序
func (this *Graph) TuoPuSort() ([]int, bool) {
	mapQianXIang := make(map[int]int)
	sortList := []int{}

	for i := 0; i < this.vertNum; i++ {
		node := this.vertNode[i]
		if node == nil {
			continue
		}

		// 获取邻接链表
		edgeNode := node.edgeTableNode
		// 记录每个子节点的被指向的次数
		for edgeNode != nil {
			mapQianXIang[edgeNode.index]++
			edgeNode = edgeNode.edgeTableNode
		}

	}

	// 拓扑排序就是把没有前项节点的先放入队列中
	for i := 0; i < this.vertNum; i++ {
		if mapQianXIang[i] == 0 {
			queue.Put(MapParent{parent: -1, son: i})
		}
	}

	for !queue.IsEmpty() {
		node := queue.Pop()

		sortList = append(sortList, node.value.son)

		// 获取邻接链表
		edgeNode := this.vertNode[node.value.son].edgeTableNode
		// 递减指向次数 其实就是斩断的过程
		for edgeNode != nil {
			mapQianXIang[edgeNode.index]--
			if mapQianXIang[edgeNode.index] == 0 {
				// 忽略parent 我们son
				queue.Put(MapParent{parent: -1, son: edgeNode.index})
			}
			edgeNode = edgeNode.edgeTableNode
		}
	}

	fmt.Println(sortList, "  ", len(sortList) != this.vertNum)
	// bool 表示有没有环
	return sortList, len(sortList) != this.vertNum
}

// 获取权重最低的节点
// 不能包含已经搜索过的
func (this *Graph) MinCostsNode(costs map[int]float64) int {
	minCostIndex := -1

	inf := math.Inf(1)
	for key, value := range costs {
		// 已经被查找过
		if this.visted[key] == 1 {
			continue
		}

		if value < inf {
			minCostIndex = key
			inf = value
		}
	}
	return minCostIndex
}

func (this *Graph) Dijkstra(start, end int) []int {
	if start >= this.vertNum || start < 0 || end >= this.vertNum || end < 0 {
		return nil
	}
	// 除起点外所有要到达的节点的对应权重
	mapCosts := make(map[int]float64)
	// 从哪个节点到这个节点的映射关系
	// key值与上面的mapCosts一样
	mapParent := make(map[int]float64)

	// 匿名函数初始化
	func() {
		//
		for i := 0; i < this.vertNum; i++ {
			if i == start {
				node := this.vertNode[i]
				edgeNode := node.edgeTableNode
				for edgeNode != nil {
					mapCosts[edgeNode.index] = float64(edgeNode.weight)
					mapParent[edgeNode.index] = float64(i)
					edgeNode = edgeNode.edgeTableNode
				}

				continue
			}

			// 其他非起点的邻居节点设置为∞
			if _, ok := mapCosts[i]; !ok {
				mapCosts[i] = math.Inf(0)
				mapParent[i] = math.Inf(0)
			}
		}
	}()

	minCostIndex := this.MinCostsNode(mapCosts)
	for minCostIndex != -1 {
		// 获取他的邻接点
		node := this.vertNode[minCostIndex]
		// 获取他的所有邻接点
		edgeNode := node.edgeTableNode
		for edgeNode != nil {
			// 如果到达邻接点的权重+最小权重的值小于邻接点的权重就更新  起点到达邻接点的权重
			tempCost := float64(edgeNode.weight) + mapCosts[minCostIndex]
			if tempCost < mapCosts[edgeNode.index] {
				mapCosts[edgeNode.index] = tempCost
				// 更新父节点
				mapParent[edgeNode.index] = float64(minCostIndex)
			}

			// 查找剩余的邻接点
			edgeNode = edgeNode.edgeTableNode
		}

		// 当前节点已经处理过了
		this.visted[minCostIndex] = 1
		minCostIndex = this.MinCostsNode(mapCosts)
	}

	var path []int

	// 逆序遍历出路径
	if math.IsInf(mapParent[end], 0) {
		return path
	}

	// 查找路径
	path = append(path, end)
	value := int(mapParent[end])
	for value != start {
		path = append(path, value)
		value = int(mapParent[value])
	}

	path = append(path, start)

	// 是逆序的寻路点
	// 可以倒序得到顺序的节点
	fmt.Println(path)
	return path
}

/**
栈：限制插入和删除只能在一个位置上进行的表，该位置是表的末端，叫做栈的顶(top)
对栈的基本操作有push(进栈)和pop(出栈)。
基本算法：
进栈(push):
1.若top>=n时，作出错误处理(进栈前先检查栈是否已满，满则溢出，不满则进入2)
2.置top = top + 1(栈指针加1,指向进栈地址)
3.s(top) = x ,结束(x为新进栈的元素)
出栈(pop):
1.若top <=0，则给出下溢信息，作出错处理(出栈前先检查栈是否为空，空则下溢，不空走2)
2.x = s(top),出栈后的元素赋值给x
3.top = top -1 ，栈指针减1,指向栈顶
*/

// 定义常量栈的初始大小
const initSize int = 20

// 栈结构体
type Stack struct {
	// 容量
	size int
	// 栈顶
	top int
	// 用slice作容器，定义为interface{}
	data []interface{}
}

// 创建并初始化栈，返回strck
func CreateStack() Stack {
	s := Stack{}
	s.size = initSize
	s.top = -1
	s.data = make([]interface{}, initSize)
	return s
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// 判断栈是否已满
func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

// 入栈
func (s *Stack) Push(data interface{}) bool {
	// 首先判断栈是否已满
	if s.IsFull() {
		fmt.Println("stack is full, push failed")
		return false
	}
	// 栈顶指针+1
	s.top++
	// 把当前的元素放在栈顶的位置
	s.data[s.top] = data
	return true
}

// 获取栈顶元素
func (s *Stack) Get() interface{} {
	// 判断是否是空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty , pop error")
		return nil
	}

	// 把栈顶的元素赋值给临时变量tmp
	tmp := s.data[s.top]
	return tmp
}

// pop,返回栈顶元素
func (s *Stack) Pop() interface{} {
	// 判断是否是空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty , pop error")
		return nil
	}
	// 把栈顶的元素赋值给临时变量tmp
	tmp := s.data[s.top]
	// 栈顶指针-1
	s.top--
	return tmp
}

// 栈的元素的长度
func (s *Stack) GetLength() int {
	length := s.top + 1
	return length
}

// 清空栈
func (s *Stack) Clear() {
	s.top = -1
}

// 遍历栈
func (s *Stack) Traverse() {
	// 是否为空栈
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}

	for i := 0; i <= s.top; i++ {
		fmt.Println(s.data[i], " ")
	}
}
