package main

import "sort"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int) //key:a+b的数值，value:a+b数值出现的次数
	count := 0
	// 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 遍历nums3和nums4数组，找到如果 0-(c+d) 在map中出现过的话，就把map中key对应的value也就是出现次数统计出来
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}
	return count
}

func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine { // 通过recode数据记录 magazine里各个字符出现次数
		record[v-'a']++
	}
	for _, v := range ransomNote { // 遍历ransomNote，在record里对应的字符个数做--操作
		record[v-'a']--
		if record[v-'a'] < 0 { // 如果小于零说明ransomNote里出现的字符，magazine没有
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
