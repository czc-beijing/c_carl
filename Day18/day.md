
# 算法记录 | Day18 二叉树


## LeetCode  513.找树左下角的值
给定一个二叉树，在树的最后一行找到最左边的值
###  解题思路
#### 
```
示例:
```

代码
```
var depth int   
var res int    
func findBottomLeftValue(root *TreeNode) int {
    depth, res = 0, 0   // 初始化
    dfs(root, 1)
    return res
}

func dfs(root *TreeNode, d int) {
    if root == nil {
        return
    }
    // 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
    if root.Left == nil && root.Right == nil && depth < d { 
        depth = d
        res = root.Val
    }
    dfs(root.Left, d+1)   // 隐藏回溯
    dfs(root.Right, d+1)
}
```


## LeetCode 112. 路径总和

###  解题思路
#### 
```
示例:
```
代码
```
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}
	return hasPathSum(root.Left, targetSum - root.Val) ||  hasPathSum(root.Right, targetSum - root.Val)
}
```




## LeetCode 113. 路径总和ii
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点
###  解题思路
#### 
```
示例:
```
代码
```
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
    path := []int{}
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, left int) {
        if node == nil {
            return
        }
        left -= node.Val
        path = append(path, node.Val)
        defer func() { path = path[:len(path)-1] }()
        if node.Left == nil && node.Right == nil && left == 0 {
            ans = append(ans, append([]int(nil), path...))
            return
        }
        dfs(node.Left, left)
        dfs(node.Right, left)
    }
    dfs(root, targetSum)
    return
}
```


## LeetCode 106.从中序与后序遍历序列构造二叉树，105.从前序与中序遍历序列构造二叉树 一起做，思路一样的
根据一棵树的中序遍历与后序遍历构造二叉树。

注意: 你可以假设树中没有重复的元素。

###  解题思路
#### 
```
例如，给出 
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3] 返回如下的二叉树
```
代码
```

```