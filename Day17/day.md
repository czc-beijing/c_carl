
# 算法记录 | Day13 二叉树


## LeetCode  110.平衡二叉树
###  解题思路
#### 
```
示例 2：

```

代码
```
func isBalanced(root *TreeNode) bool {
    if getHeight(root) < 0{
        return  false
    }
     return  true
}

func getHeight(root *TreeNode) int {
    if root == nil{
        return 0
    }
    l,r := getHeight(root.Left), getHeight(root.Right)
    if l == -1 || r == -1{
        return -1
    }
    if l - r > 1 || r - l > 1{
        return -1
    }
    return max(l,r) + 1
}

func max(i,j int) int{
    if i > j {
        return i
    }
    return j
}
```





## LeetCode 257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点
###  解题思路
#### 
```
示例:
```
代码
```
func isBalanced(root *TreeNode) bool {
    if getHeight(root) < 0{
        return  false
    }
     return  true
}

func getHeight(root *TreeNode) int {
    if root == nil{
        return 0
    }
    l,r := getHeight(root.Left), getHeight(root.Right)
    if l == -1 || r == -1{
        return -1
    }
    if l - r > 1 || r - l > 1{
        return -1
    }
    return max(l,r) + 1
}

func max(i,j int) int{
    if i > j {
        return i
    }
    return j
}
```




## LeetCode 257. 二叉树的所有路径
力扣题目链接(opens new window)
给定一个二叉树，返回所有从根节点到叶子节点的路径。
###  解题思路
#### 
```
示例:
```
代码
```
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
```



## LeetCode 404. 左叶子之和
给定二叉树的根节点 root ，返回所有左叶子之和。
###  解题思路
#### 
```
示例:
```
代码
```
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0;
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
         res = root.Left.Val        // 中
    }
    return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) +  res
}
```