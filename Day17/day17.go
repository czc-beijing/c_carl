package main

import (
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	pathJoin(root, "")
	return res
}

func pathJoin(node *TreeNode, s string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		v := s + strconv.Itoa(node.Val)
		res = append(res, v)
		return
	}
	s = s + strconv.Itoa(node.Val) + "->"
	if node.Left != nil {
		pathJoin(node.Left, s)
	}
	if node.Right != nil {
		pathJoin(node.Right, s)
	}
	return
}
