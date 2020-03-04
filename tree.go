package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
