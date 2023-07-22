# 算法记录 | Day2 数组基础


## LeetCode 209. 长度最小的子数组
### 题目
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回

示例 1:
```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
```
示例 2:
```
输入：target = 4, nums = [1,4,4]
输出：1
 ```

###  解题思路
- 滑动窗口

#### 
```
func minSubArrayLen(target int, nums []int) int {
	i := 0
	ans := math.MaxInt32
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			ans = min(ans, j-i+1)
			sum = sum - nums[i]
			i++
		}
	}
    if ans == math.MaxInt32 {
        return 0
    }
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

```

###  解题思路


## LeetCode 209. 长度最小的子数组
### 题目
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回

示例 1:
```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
```
示例 2:
```
输入：target = 4, nums = [1,4,4]
输出：1
 ```

###  解题思路
- 滑动窗口

#### 
```
func minSubArrayLen(target int, nums []int) int {
	i := 0
	ans := math.MaxInt32
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			ans = min(ans, j-i+1)
			sum = sum - nums[i]
			i++
		}
	}
    if ans == math.MaxInt32 {
        return 0
    }
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

```

###  解题思路




## LeetCode 209. 螺旋矩阵 II
### 题目
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix


示例 1:

```
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
示例 2：

```
示例 2:
```

输入：n = 1
输出：[[1]]
 ```

###  解题思路
-  a[0][0], a[1][1]  offset loop  startX, startY

#### 
```

func generateMatrix(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	startX, startY := 0, 0
	loop := n / 2
	center := n / 2
	offset := 1
	count := 1
	for loop > 0 {
		i, j := startX, startY
		for j = startY; j < n-offset; j++ {
			res[i][j] = count
			count++
		}
		for i = startX; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}
		offset++
		loop--
		startX++
		startY++
	}
	if n%2 == 1 {
		res[center][center] = count
	}
	return res
}

```

###  解题思路



##  总结
