
# 算法记录 | Day8 字符串


## LeetCode 28. 实现 strStr()
### 题目
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符
示例:
```
示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2

示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1

```


###  解题思路
-  暴力破解
-

#### 
```
class Solution {
    public int strStr(String ss, String pp) {
        int n = ss.length(), m = pp.length();
        char[] s = ss.toCharArray(), p = pp.toCharArray();
        // 枚举原串的「发起点」
        for (int i = 0; i <= n - m; i++) {
            // 从原串的「发起点」和匹配串的「首位」开始，尝试匹配
            int a = i, b = 0;
            while (b < m && s[a] == p[b]) {
                a++;
                b++;
            }
            // 如果能够完全匹配，返回原串的「发起点」下标
            if (b == m) return i;
        }
        return -1;
    }
}

```

###  解题思路
## LeetCode   459.重复的子字符串
### 题目
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例：
```
示例 1:

输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"
输出: False
示例 3:

输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
```
#### 
```
func repeatedSubstringPattern(s string) bool {
    return kmp(s)
}

func kmp(pattern string) bool {
    n := len(pattern)
    fail := make([]int, n)
    for i := 0; i < n; i++ {
        fail[i] = -1
    }
    for i := 1; i < n; i++ {
        j := fail[i - 1]
        for (j != -1 && pattern[j + 1] != pattern[i]) {
            j = fail[j]
        }
        if pattern[j + 1] == pattern[i] {
            fail[i] = j + 1
        }
    }
    return fail[n - 1] != -1 && n % (n - fail[n - 1] - 1) == 0
}