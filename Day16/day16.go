package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	que := make([]*TreeNode, 0)
	result := make([][]int, 0)
	que = append(que, root)
	for len(que) > 0 {
		size := len(que)
		item := make([]int, 0)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			item = append(item, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := 0
	rightDepth := 0
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}
	if leftDepth < rightDepth {
		leftDepth = rightDepth
	}
	return leftDepth + 1
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
