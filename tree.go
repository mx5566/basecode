package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode

	Parent *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (this *Tree) GetRoot() *TreeNode {
	return this.root
}

// 中序 遍历创建二叉树 通过数组
func BuildTreeMidSeq(arr []int) *Tree {
	var tree = Tree{root: nil}

	var length = len(arr)
	if length == 0 {
		return &tree
	}

	rootNode := new(TreeNode)
	rootNode.Val = arr[0]
	rootNode.Left = nil
	rootNode.Right = nil

	tree.root = rootNode

	currentNode := new(TreeNode)
	currentNode.Left = rootNode
	currentNode.Right = rootNode

	var createTree = func(root *TreeNode, i int) {
		root.Val = arr[i]
		root.Left = nil
		root.Right = nil
	}

	for i := 1; i < length; i += 2 {
		if arr[i] != int(nil) {
			currentNode.Left.Left = new(TreeNode)
			createTree(rootNode.Left.Left, i)

			currentNode.Left = currentNode.Left.Left

		}

		if i+1 < length && arr[i+1] != int(nil) {
			currentNode.Right.Right = new(TreeNode)
			createTree(currentNode.Right.Right, i+1)

			currentNode.Right = currentNode.Right.Right
		}

	}

	return &tree
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth = 0
	var maxDepth = 0
	dfsDepth(root, depth, &maxDepth)
	return maxDepth
}

// 通过DFS求二叉树的最大深度
func dfsDepth(root *TreeNode, depth int, maxDepth *int) {
	if root == nil {
		return
	}

	depth += 1

	if root.Left != nil {
		dfsDepth(root.Left, depth, maxDepth)
	}

	if root.Right != nil {
		dfsDepth(root.Right, depth, maxDepth)
	}

	if root.Left == nil && root.Right == nil {
		if depth > *maxDepth {
			*maxDepth = depth
		}
	}
}

// 通过BFS求二叉树的最大深度
func bfsDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	queue = append(queue, root)

	depth := 0
	for len(queue) > 0 {
		depth += 1

		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[length:]
	}

	return depth
}

func serachTreedfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		serachTreedfs(root.Left)
	}

	fmt.Println(root.Val)

	if root.Right != nil {
		serachTreedfs(root.Right)
	}

}

// 使用栈作为辅助结构，采用中序遍历二叉搜索树
//
func searchTreeStack(root *TreeNode) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	var rootTemp = root

	for rootTemp != nil || len(stack) > 0 {
		for rootTemp != nil {
			stack = append(stack, rootTemp)
			rootTemp = rootTemp.Left
		}

		length := len(stack)
		rootTemp := stack[length-1]
		fmt.Println(rootTemp.Val)
		rootTemp = rootTemp.Right
	}
}

// 验证是不是二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root)

	var minS = []int{-1 << 63}
	var maxS = []int{1<<63 - 1}

	for len(queue) > 0 {
		node := queue[0]

		min := minS[0]
		max := maxS[0]

		if min >= node.Val || max <= node.Val {
			return false
		}

		queue = queue[1:]
		minS = minS[1:]
		maxS = maxS[1:]
		if node.Left != nil {
			minS = append(minS, min)
			maxS = append(maxS, node.Val)
			queue = append(queue, node.Left)

		}

		if node.Right != nil {
			minS = append(minS, node.Val)
			maxS = append(maxS, max)
			queue = append(queue, node.Right)
		}
	}

	return true
}

func insertBST(root *TreeNode, value int) {
	if root == nil {
		node := new(TreeNode)
		node.Right, node.Left = nil, nil
		node.Val = value

		root = node
		return
	}

	if root.Val <= value {
		if root.Right == nil {
			node := new(TreeNode)
			node.Right, node.Left = nil, nil
			node.Val = value

			root.Right = node
		} else {
			insertBST(root.Right, value)
		}
	} else {
		if root.Left == nil {
			node := new(TreeNode)
			node.Right, node.Left = nil, nil
			node.Val = value

			root.Left = node
		} else {
			insertBST(root.Left, value)
		}
	}
}

func insertBSTWhile(tree *Tree, value int) {
	node := new(TreeNode)
	node.Left, node.Right = nil, nil
	if tree.root == nil {
		tree.root = node
		return
	}

	root := tree.root
	var findNode *TreeNode = nil
	for root != nil {
		findNode = root
		if root.Val <= value {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	if findNode == nil {
		return
	}

	if findNode.Val <= value {
		findNode.Right = node
	} else {
		findNode.Left = node
	}

}

func FindBSTMinNum(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 沿着左边走
	for root.Left != nil {
		root = root.Left
	}

	return root
}

func Rorate(tree *Tree, deleteNode *TreeNode, replaceNode *TreeNode) {
	// 没有父节点 肯定是根节点
	if deleteNode.Parent == nil {
		tree.root = replaceNode
		// 说明这个节点是其父节点 的左节点
	} else if deleteNode.Parent.Left == deleteNode {
		deleteNode.Parent.Left = replaceNode

		// 说明这个节点是其父节点 的左节点
	} else if deleteNode.Parent.Right == deleteNode {
		deleteNode.Parent.Right = replaceNode
	}

	// 更新父节点
	if replaceNode != nil {
		replaceNode.Parent = deleteNode.Parent
	}
}

func DeleteBST(tree *Tree, deleteNode *TreeNode) {
	if tree == nil {
		return
	}

	if tree.root == nil {
		return
	}

	if deleteNode == nil {
		return
	}

	// 如果删除的节点没有左节点
	if deleteNode.Left == nil {
		Rorate(tree, deleteNode, deleteNode.Right)
	} else if deleteNode.Right == nil {
		Rorate(tree, deleteNode, deleteNode.Left)
	} else {
		// 找到他的右节点的后继节点
		minNumNode := FindBSTMinNum(deleteNode.Right)
		if minNumNode == nil {
			return
		}

		if minNumNode.Parent != deleteNode {
			Rorate(tree, minNumNode, minNumNode.Right)
			minNumNode.Right = deleteNode
			minNumNode.Right.Parent = minNumNode
		}

		Rorate(tree, deleteNode, minNumNode)
		minNumNode.Left = deleteNode.Left
		minNumNode.Left.Parent = minNumNode
	}
}

const (
	none  = 0
	left  = 1
	right = 2
)

type DireQueue struct {
	val       int
	direction int
}

func InordertreeWalk(root *TreeNode, queue *[]DireQueue, direction int) {
	if root == nil {
		return
	}

	InordertreeWalk(root.Left, queue, left)

	*queue = append(*queue, DireQueue{val: root.Val, direction: direction})
	println(root.Val)

	InordertreeWalk(root.Right, queue, right)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归的方案
	return isMirror(root.Left, root.Right)
}

func bfsMirror(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root)
	queue = append(queue, root)

	for len(queue) > 0 {
		t1 := queue[0]
		t2 := queue[1]

		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}

		if t1 == t2 {
			queue = append(queue, t1.Left)
			queue = append(queue, t1.Right)
		} else {
			queue = append(queue, t1.Left)
			queue = append(queue, t2.Right)

			queue = append(queue, t1.Right)
			queue = append(queue, t2.Left)
		}
	}

	return true
}
