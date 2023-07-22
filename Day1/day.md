# 算法记录 | Day1数组基础

## LeetCode 704-二分查找

### 题目
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
```
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
```
示例 2:
```
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
 ```

###  解题思路
- 快慢指针
- 循环

#### 二分法-第一种写法（左闭右闭区间 [left, right] ）
```
func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l<=r {
        mid := (r-l) >> 1 + l
        if nums[mid] == target{
            return mid
        }
        if nums[mid] > target{
            r = mid-1
        }
         if nums[mid] < target{
            l = mid+1
        }
    }
    return -1
}
```

###  解题思路
- 假设在 nums = [0, 1, 2, 3, 4, 5] 中查找3。由于是左闭右闭区间，需要包含 0 与 5 -->[0, 5]，因此 left 指针为 0，right 指针为 5， 即为len(nums) - 1.
- 在 while 循环中，若查找到 [3, 3]，left == right 是可以满足区间定义的，因此 while left <= right。当进行if else的判断的时候，mid就会已经被判断了，此时直接都进行 +1 或者 -1的操作即可。



## LeetCode 27-移除元素
### 题目
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
```
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
```
示例 2:
```
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
 ```

###  解题思路
- 快慢指针

#### 
```
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast <= len(nums)-1; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow += 1
		}
	}
	return slow
}
```

###  解题思路
- 假设在 nums = [0, 1, 2, 3, 4, 5] 中查找3。由于是左闭右闭区间，需要包含 0 与 5 -->[0, 5]，因此 left 指针为 0，right 指针为 5， 即为len(nums) - 1.
- 在 while 循环中，若查找到 [3, 3]，left == right 是可以满足区间定义的，因此 while left <= right。当进行if else的判断的时候，mid就会已经被判断了，此时直接都进行 +1 或者 -1的操作即可。

##  总结