
# 算法记录 | Day11 堆栈


## LeetCode 239. 滑动窗口最大值
### 题目
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值
```
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 
示例 2：
输入：nums = [1], k = 1
输出：[1]
```


###  解题思路
-  栈
-

#### 
```

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}
```

###  解题思路
## LeetCode  347.前 K 个高频元素
### 题目
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
#### 

示例
````
示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
````
实现
```
func topKFrequent(nums []int, k int) []int {
    ans:=[]int{}
    map_num:=map[int]int{}
    for _,item:=range nums {
        map_num[item]++
    }
    for key,_:=range map_num{
        ans=append(ans,key)
    }
    //核心思想：排序
    //可以不用包函数，自己实现快排
    sort.Slice(ans,func (a,b int)bool{
        return map_num[ans[a]]>map_num[ans[b]]
    })
    return ans[:k]
}
```

## LeetCode   150. 逆波兰表达式求值
### 题目
根据 逆波兰表示法，求表达式的值。
有效的运算符包括 + ,  - ,  * ,  / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
说明：
整数除法只保留整数部分。 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
#### 

示例
````
示例 1：

输入: ["2", "1", "+", "3", " * "]
输出: 9
解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: 该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", " * ", "/", " * ", "17", "+", "5", "+"]

输出: 22
````
实现
```
func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		m, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, m)
			continue
		}
		m1 := stack[len(stack)-1]
		m2 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if tokens[i] == "+" {
			stack = append(stack, m2+m1)
		}
		if tokens[i] == "-" {
			stack = append(stack, m2-m1)
		}
		if tokens[i] == "*" {
			stack = append(stack, m2*m1)
		}
		if tokens[i] == "/" {
			stack = append(stack, m2/m1)
		}
	}
	return stack[len(stack)-1]
}
```