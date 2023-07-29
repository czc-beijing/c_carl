
# 算法记录 | Day13 二叉树


## LeetCode 222.完全二叉树的节点个数
###  解题思路
#### 
```
示例 2：

```

代码
```
func countNodes(root *TreeNode) int {
     if root ==  nil{
         return 0
     }
     return countNodes(root.Left) + countNodes(root.Right) + 1
}
```



