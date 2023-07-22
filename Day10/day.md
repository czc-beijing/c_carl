
# 算法记录 | Day8 字符串


## LeetCode  232.用栈实现队列
### 题目
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。
示例:
```
MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false

```


###  解题思路
-  
-

#### 
```
type MyQueue struct {
    stackIn  []int //输入栈
    stackOut []int //输出栈
}

func Constructor() MyQueue {
    return MyQueue{
        stackIn:  make([]int, 0),
        stackOut: make([]int, 0),
    }
}

// 往输入栈做push
func (this *MyQueue) Push(x int) {
    this.stackIn = append(this.stackIn, x)
}

// 在输出栈做pop，pop时如果输出栈数据为空，需要将输入栈全部数据导入，如果非空，则可直接使用
func (this *MyQueue) Pop() int {
    inLen, outLen := len(this.stackIn), len(this.stackOut)
    if outLen == 0 {
        if inLen == 0 {
            return -1
        }
        for i := inLen - 1; i >= 0; i-- {
            this.stackOut = append(this.stackOut, this.stackIn[i])
        }
        this.stackIn = []int{}      //导出后清空
        outLen = len(this.stackOut) //更新长度值
    }
    val := this.stackOut[outLen-1]
    this.stackOut = this.stackOut[:outLen-1]
    return val
}

func (this *MyQueue) Peek() int {
    val := this.Pop()
    if val == -1 {
        return -1
    }
    this.stackOut = append(this.stackOut, val)
    return val
}

func (this *MyQueue) Empty() bool {
    return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
```

###  解题思路
## LeetCode    225. 用队列实现栈
### 题目
push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）

#### 
```
type MyStack struct {
    //创建两个队列
    queue1 []int
    queue2 []int
}


func Constructor() MyStack {
    return MyStack{	//初始化
        queue1:make([]int,0),
        queue2:make([]int,0),
    }
}


func (this *MyStack) Push(x int)  {
     //先将数据存在queue2中
    this.queue2 = append(this.queue2,x)	
   //将queue1中所有元素移到queue2中，再将两个队列进行交换
    this.Move() 
}


func (this *MyStack) Move(){    
    if len(this.queue1) == 0{
        //交换，queue1置为queue2,queue2置为空
        this.queue1,this.queue2 = this.queue2,this.queue1
    }else{
        //queue1元素从头开始一个一个追加到queue2中
            this.queue2 = append(this.queue2,this.queue1[0])  
            this.queue1 = this.queue1[1:]	//去除第一个元素
            this.Move()     //重复
    }
}

func (this *MyStack) Pop() int {
    val := this.queue1[0]
    this.queue1 = this.queue1[1:]	//去除第一个元素
    return val

}


func (this *MyStack) Top() int {
    return this.queue1[0]	//直接返回
}


func (this *MyStack) Empty() bool {
return len(this.queue1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */