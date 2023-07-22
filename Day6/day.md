# 算法记录 | Day6 哈希表


## LeetCode 42. 有效的字母异位词
### 题目

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:
```
输入: s = "anagram", t = "nagaram"
输出: true
```
示例 2:
```
输入: s = "rat", t = "car"
输出: false
 ```

###  解题思路
-  map
-

#### 
```
func isAnagram(s string, t string) bool {
    cnt := map[rune]int{}
    if len(s) != len(t){
        return false
    }
    for _, ch := range s{
        cnt[ch]++
    }
    for _, ch := range t{
        cnt[ch]--
        if cnt[ch] < 0{
            return false
        }
    }
    return true
}
```

###  解题思路
## LeetCode  349. 两个数组的交集
### 题目

题意：给定两个数组，编写一个函数来计算它们的交集

示例 1:
```
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

```
示例 2:
```
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
 ```


###  解题思路
- 

#### 
```
func intersection(nums1 []int, nums2 []int) []int {

	map1 := map[int]struct{}{}
	for _, v := range nums1 {
		map1[v] = struct{}{}
	}
	map2 := map[int]struct{}{}
	for _, num := range nums2 {
		map2[num] = struct{}{}
	}
	if len(map1) > len(map2) {
		map1, map2 = map2, map1
	}
	var result []int
	for k, _ := range map1 {
		if _, ok := map2[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

```
###  解题思路




## LeetCode 202. 快乐数
### 题目
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

示例 1:
```
输入：19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
```
###  解题思路
-  
#### 
```

func isHappy(n int) bool {
    m := make(map[int]bool)
    for n != 1 && !m[n] {
        n, m[n] = getSum(n), true
    }
    return n == 1
}

func getSum(n int) int {
    sum := 0
    for n > 0 {
        sum += (n % 10) * (n % 10)
        n = n / 10
    }
    return sum
}

```

###  解题思路



## LeetCode 1. 两数之和
### 题目

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍


示例 1:

```
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```
###  解题思路
-  
#### 
```
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, j := range nums{
        if l, ok := m[target-j];ok{
            return []int{l,i}
        }
        m[j] = i
    }
    return []int{}
}
```

###  解题思路


##  总结
