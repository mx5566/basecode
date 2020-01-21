package main

import "fmt"

type Key int32
type Value int32

//------------------------------linklist---------------------
// LinkList implements
type Node struct {
	pre   *Node
	next  *Node
	key   Key
	value Value
}

type LinkList struct {
	head *Node
	tail *Node
}

func InitLinkList() *LinkList {
	var linkList LinkList
	linkList.head = new(Node)
	linkList.tail = new(Node)

	linkList.head.next = linkList.tail
	linkList.head.pre = nil
	linkList.tail.pre = linkList.head
	linkList.tail.next = nil

	return &linkList
}

// insert first
func (this *LinkList) Insert(node *Node) {
	if node == nil {
		return
	}

	node.next = this.head.next
	node.pre = this.head

	this.head.next.pre = node
	this.head.next = node
}

func (this *LinkList) GetTail() *Node {
	return this.tail
}

// remove last one
func (this *LinkList) Remove(node *Node) {
	if nil == node {
		return
	}

	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LinkList) Print() {

	s := this.head.next
	for s != this.tail {
		fmt.Println(s.key, " ", s.value)
		s = s.next
	}
}

//------------------------------lru-------------------------

// LRU Implements
type Lru struct {
	mapList  map[Key]*Node
	maxCap   int
	linkList *LinkList
}

func (this *Lru) Put(key Key, value Value) {
	ret := this.Get(key)
	if ret == nil {
		full := this.IsFull()
		if full {
			fmt.Println("lru full")
			tail := this.linkList.GetTail()
			this.Remove(tail.pre.key)
		}

		node := new(Node)
		node.key = key
		node.value = value

		this.linkList.Insert(node)
		this.mapList[key] = node
		return
	}

	this.linkList.Remove(ret)
	this.linkList.Insert(ret)
}

func (this *Lru) Get(key Key) *Node {
	ret, ok := this.mapList[key]
	if !ok {
		return nil
	}

	return ret
}

func (this *Lru) IsFull() bool {
	return this.maxCap <= len(this.mapList)
}

func (this *Lru) Remove(key Key) {
	ret := this.Get(key)
	if ret == nil {
		return
	}

	this.linkList.Remove(ret)
	delete(this.mapList, key)
}

func (this *Lru) Print() {
	this.linkList.Print()
}

func NewLru(size int) *Lru {
	if size < 1 {
		return nil
	}
	link := InitLinkList()
	return &Lru{maxCap: size,
		linkList: link,
		mapList:  make(map[Key]*Node)}
}

func main() {
	lru := NewLru(3)
	lru.Put(1, 20)
	lru.Put(2, 21)
	lru.Put(3, 22)
	lru.Put(4, 23)
	lru.Print()

	lru.Put(5, 24)

	lru.Print()

}
