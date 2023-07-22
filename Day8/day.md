
# 算法记录 | Day8 字符串


## LeetCode 344.反转字符串
### 题目
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符
示例 1:
```
示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
```


###  解题思路
-  map
-

#### 
```
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

```

###  解题思路
## LeetCode   541. 反转字符串II
### 题目
给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样

示例 1：
```
输入: s = "abcdefg", k = 2
输出: "bacdfeg"
```
#### 
```
func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		// 1. 每隔 2k 个字符的前 k 个字符进行反转
		// 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
		if i + k <= length {
			reverse(ss[i:i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
```
###  解题思路
## LeetCode 剑指Offer 05.替换空格
### 题目

示例 1:
```
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1： 输入：s = "We are happy."
输出："We%20are%20happy."
```
###  解题思路
-  
#### 
```
// 遍历添加
func replaceSpace(s string) string {
    b := []byte(s)
    result := make([]byte, 0)
    for i := 0; i < len(b); i++ {
        if b[i] == ' ' {
            result = append(result, []byte("%20")...)
        } else {
            result = append(result, b[i])
        }
    }
    return string(result)
}

```
## LeetCode 151.翻转字符串里的单词
### 题目
给定一个字符串，逐个翻转字符串中的每个单词。
#
示例:
```
示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个
```
###  解题思路
-  
#### 
```
func reverseWords(s string) string {
    b := []byte(s)

    // 移除前面、中间、后面存在的多余空格
    slow := 0
    for i := 0; i < len(b); i++ {
        if b[i] != ' ' {
            if slow != 0 {
                b[slow] = ' '
                slow++
            }
            for i < len(b) && b[i] != ' ' { // 复制逻辑
                b[slow] = b[i]
                slow++
                i++
            }
        }
    }
    b = b[0:slow]
    
    // 翻转整个字符串
    reverse(b)
    // 翻转每个单词
    last := 0
    for i := 0; i <= len(b); i++ {
        if i == len(b) || b[i] == ' ' {
            reverse(b[last:i])
            last = i + 1
        }
    }
    return string(b)
}

func reverse(b []byte) {
    left := 0
    right := len(b) - 1
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
}
```

## LeetCode  剑指Offer58-II.左旋转字符串
### 题目
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"
#
示例:
```
示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"
```
###  解题思路
-  
#### 
```
func reverseLeftWords(s string, n int) string {
    b := []byte(s)
    // 1. 反转前n个字符
    // 2. 反转第n到end字符
    // 3. 反转整个字符
    reverse(b, 0, n-1)
    reverse(b, n, len(b)-1)
    reverse(b, 0, len(b)-1)
    return string(b)
}
// 切片是引用传递
func reverse(b []byte, left, right int){
    for left < right{
        b[left], b[right] = b[right],b[left]
        left++
        right--
    }
}
```
