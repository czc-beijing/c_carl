# 算法记录 | Day3 链表


## LeetCode 203. 移除链表元素
### 题目
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1:
```
输入：head = [], val = 1
输出：[]
```
示例 2:
```
输入：head = [7,7,7,7], val = 7
输出：[]
 ```

###  解题思路
- 临时节点
- next节点

#### 
```
func removeElements(head *ListNode, val int) *ListNode {
	node := &ListNode{
		Val :0,
		Next: head,
    }
    var tmp = node
	for node.Next != nil{
		if node.Next.Val == val{
             node.Next = node.Next.Next
         }else{
			 node = node.Next
		 }
    }
	return tmp.Next
}
```

###  解题思路


## LeetCode 707. 设计链表
### 题目
你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始

示例 1:
```
输入
["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
[[], [1], [3], [1, 2], [1], [1], [1]]
输出
[null, null, null, null, 2, null, 3]

解释
MyLinkedList myLinkedList = new MyLinkedList();
myLinkedList.addAtHead(1);
myLinkedList.addAtTail(3);
myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
myLinkedList.get(1);              // 返回 2
myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
myLinkedList.get(1);              // 返回 3
 ```

###  解题思路
- 

#### 
```
/*
这道题要搞清楚head节点和第一个节点的区别
head节点其实是第0号节点
*/

type Node struct {
	val  int   //链表的值
	next *Node //指向下一个节点的指针
}

type MyLinkedList struct {
	head *Node //虚拟头结点
	size int   //定义链表的长度
}

// 对链表的初始化操作
func Constructor() MyLinkedList {
	// 将head节点的值设置为-1,指向空节点
	dummyHead := &Node{
		val:  -1,
		next: nil,
	}
	// 将链表的的长度设置为1
	return MyLinkedList{head: dummyHead, size: 0}
}

// 通过index下标获取对应的val
func (this *MyLinkedList) Get(index int) int {
	// 判断index是否合法
	if (index < 0) || (index > (this.size - 1)) {
		return -1
	}
	// 查找操作
	var cur *Node = this.head //从head节点开始遍历
	for i := 0; i <= index; i++ {
		cur = cur.next //第index节点其实是index+1
	}
	return cur.val
}

// 将节点插入第一个节点之前,head节点之后
func (this *MyLinkedList) AddAtHead(Val int) {
	// 先创建一个新的节点
	var newNode *Node = &Node{val: Val, next: nil}
	// 将新节点指向第一个节点
	newNode.next = this.head.next
	// 然后在将头结点指向新的节点
	this.head.next = newNode
	// 然后将链表的长度加1
	this.size++
}

// 将节点插入到链表的最后面
func (this *MyLinkedList) AddAtTail(Val int) {
	// 先创建一个新的节点
	var newNode *Node = &Node{val: Val, next: nil}
	// 创建一个当前节点，并将它指向head结点，可以说它就是head结点
	cur := this.head
	// 因为是插到结尾
	for cur.next != nil {
		cur = cur.next
	}
	// 此时的cur指向的是nil, 将新节点插入cur之后
	cur.next = newNode
	this.size++
}

// 将节点插入到任意位置
func (this *MyLinkedList) AddAtIndex(index int, Val int) {
	// 条件判断:题目要求
	// 添加的下标大于链表的长度，无效
	if index > this.size {
		return
	} else if index == this.size { //如果 index 等于链表的长度，则该节点将附加到链表的末尾
		this.AddAtTail(Val)
		return
	} else if index < 0 { //如果index小于0，则在头部插入节点。
		index = 0
	}
	// 定义一个新节点
	var newNode *Node = &Node{val: Val, next: nil}
	// 定义一个当前节点,指向头结点
	cur := this.head
	// 在index节点之前index-1处停止
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	// cur.next->index
	newNode.next = cur.next
	cur.next = newNode
	this.size++
}

// 删除节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.size-1 || index < 0 {
		return
	}
	// 从head节点开始
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	// 将当前节点指向自已节点下一个节点的下一个节点
	cur.next = cur.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */


```

###  解题思路




## LeetCode 206.反转链表
### 题目
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。


示例 1:

```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

```
示例 2:
```
输入：head = [1,2]
输出：[2,1]
 ```

###  解题思路
-  前节点，当前节点，后节点
#### 
```
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    var next *ListNode

    for head != nil {
        next = head.Next
        head.Next = prev


        prev = head
        head = next
    }
    return prev
}


```

###  解题思路


##  总结
