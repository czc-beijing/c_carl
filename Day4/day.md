# 算法记录 | Day4 链表


## LeetCode 两两交换链表中的节点
### 题目
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
```
输入：head = [1,2,3,4]
输出：[2,1,4,3]
```
示例 2:
```
输入：head = []
输出：[]
 ```

###  解题思路
-  tmp tmp1
-  三个步骤

#### 
```
func swapPairs(head *ListNode) *ListNode {
    dummyHead := &ListNode{
        Val:0, 
        Next:head,
    }
    curr := dummyHead
    for curr.Next != nil && curr.Next.Next != nil{
        tmp1 := curr.Next
        tmp2 := curr.Next.Next.Next

        curr.Next = curr.Next.Next
        curr.Next.Next = tmp1
        tmp1.Next = tmp2

        curr = curr.Next.Next
    }

    return dummyHead.Next
}
```

###  解题思路


## LeetCode  19.删除链表的倒数第N个节点
### 题目

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

###  解题思路
- 

#### 
```
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{
        Val:0,
        Next:head,
    }
    fast := dummyHead
    slow := dummyHead
    for i:=0;i<n;i++{
        fast = fast.Next
    }
    for fast.Next != nil{
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummyHead.Next
}

```
###  解题思路




## LeetCode 02.07. 链表相交
### 题目
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null

示例 1:

```
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
```
###  解题思路
-  
#### 
```

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    curA := headA
    curB := headB
    lenA, lenB := 0, 0
    step := 0
    var fast, slow *ListNode
    for curA != nil {
        curA = curA.Next
        lenA++
    }
    for curB != nil {
        curB = curB.Next
        lenB++
    }
    if lenA > lenB{
        step = lenA - lenB
        fast = headA
        slow = headB
    }else{
        step = lenB - lenA
        fast = headB
        slow = headA
    }
    for i:=0; i < step; i++ {
        fast = fast.Next
    }
    for fast != slow{
        fast = fast.Next
        slow = slow.Next
    }
    return fast
}

```

###  解题思路



## LeetCode 142.环形链表II
### 题目
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况


示例 1:

```
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
```
###  解题思路
-  
#### 
```
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            for slow != head {
                slow = slow.Next
                head = head.Next
            }
            return head
        }
    }
    return nil
}
```

###  解题思路


##  总结
