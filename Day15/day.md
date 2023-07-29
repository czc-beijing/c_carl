
# 算法记录 | Day13 二叉树


## LeetCode 144. 二叉树的层序遍历
###  解题思路
#### 
```
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
```

代码
```
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

```




## LeetCode 144. 二叉树的层序遍历II
###  解题思路
#### 
```
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
```

代码
```
func levelOrderBottom(root *TreeNode) [][]int {
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
    for i:=0;i<len(result)/2;i++{
        result[i], result[len(result)-i-1] =   result[len(result)-i-1], result[i]
    }
    return result
}
```





## LeetCode 199. 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
###  解题思路
#### 
```
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
```

代码
```
func rightSideView(root *TreeNode) []int {
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
			if i == size -1{
			item = append(item, node.Val)
			}
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
```






## LeetCode 637.二叉树的层平均值
###  本题就是层序遍历的时候把一层求个总和在取一个均值
#### 
```
示例 2：
```
代码
```
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
		return []float64{}
	}
	que := make([]*TreeNode, 0)
	result := make([]float64, 0)
	que = append(que, root)
	for len(que) > 0 {
		size := len(que)
        var sum int
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
            sum += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		 result = append(result, float64(sum)/float64(size))
	}
    return re
```




## LeetCode 104 二叉树的最大深度
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
###  解题思路
#### 
```
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
```

代码
```
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
```



## LeetCode 111 二叉树的最小深度
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
###  解题思路
#### 
```
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
```

代码
```

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
```






## LeetCode 226. 翻转二叉树 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
###  解题思路
#### 
```
示例 2：
```

代码
```
func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return root
    }
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

```


## LeetCode 101. 对称二叉树 （优先掌握递归）
给定一个二叉树，检查它是否是镜像对称的
###  解题思路
#### 
```
示例 2：
```

代码
```
func isSymmetric(root *TreeNode) bool {
    return compre(root.Left, root.Right)
}

func compre(left *TreeNode, right *TreeNode) bool{
     if left == nil && right == nil {
        return true;
    };
    if left == nil || right == nil {
        return false;
    };
    if left.Val != right.Val {
        return false;
    }
    return compre(left.Left, right.Right) &&  compre(left.Right, right.Left)
}
```

