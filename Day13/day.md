
# 算法记录 | Day13 二叉树


## LeetCode 144. 二叉树遍历
###  解题思路
#### 
```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		vals = append(vals, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		preorder(node.Right)
		vals = append(vals, node.Val)
	}
	preorder(root)
	return
}
```
### 迭代遍历
```
前序遍历
func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
	if root == nil {
		return ans
	}
	st := list.New()
    st.PushBack(root)
    for st.Len() > 0 {
        node := st.Remove(st.Back()).(*TreeNode)
        ans = append(ans, node.Val)
        if node.Right != nil {
            st.PushBack(node.Right)
        }
        if node.Left != nil {
            st.PushBack(node.Left)
        }
    }
    return ans
}

后续遍历
func postorderTraversal(root *TreeNode) []int {
    ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
    st.PushBack(root)

    for st.Len() > 0 {
        node := st.Remove(st.Back()).(*TreeNode)

        ans = append(ans, node.Val)
        if node.Left != nil {
            st.PushBack(node.Left)
        }
        if node.Right != nil {
            st.PushBack(node.Right)
        }
    }
    reverse(ans)
    return ans
}

func reverse(a []int) {
    l, r := 0, len(a) - 1
    for l < r {
        a[l], a[r] = a[r], a[l]
        l, r = l+1, r-1
    }
}

中序遍历
func inorderTraversal(root *TreeNode) []int {
    ans := []int{}
    if root == nil {
        return ans
    }

    st := list.New()
    cur := root

    for cur != nil || st.Len() > 0 {
        if cur != nil {
            st.PushBack(cur)
            cur = cur.Left
        } else {
            cur = st.Remove(st.Back()).(*TreeNode)
            ans = append(ans, cur.Val)
            cur = cur.Right
        }
    }

    return ans
}
```

