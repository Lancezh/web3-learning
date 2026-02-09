package homework01

import (
	"sort"
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 定义个 map 统计次数，最后找出 val = 1 的即可
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// 转成字符串，然后反转后按位比较
	// 也可以 前后双指针比较处理

	// 负数不算回文
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// 栈处理，遇到左括号入栈，遇到右括号比较前一位，如果相同出栈，不同直接返回 false，最后栈为空则有效
	stack := make([]byte, 0)

	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != match[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// 取最小公共集合，以第一个数组作为基准，依次比较，取后续最短的前缀，直到最后
	// 先判断是否为空
	length := len(strs)
	if length > 0 {
		// 以第一个数组作为基准
		prefix := strs[0]
		currentPrefixLen := len(prefix)

		// 遍历后续数组，比较前缀，取最短的前缀
		for i := 1; i < length; i++ {
			item := strs[i]

			// 取最短长度
			if len(item) < currentPrefixLen {
				currentPrefixLen = len(item)
			}

			for j := 0; j < currentPrefixLen; j++ {
				if item[j] != prefix[j] {
					currentPrefixLen = j
					break
				}
			}

			if currentPrefixLen == 0 {
				return ""
			}
			prefix = prefix[:currentPrefixLen]
		}
		return prefix
	}

	return ""
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 数组最后一位加一，超过 10，进位，往前遍历增加，循环到头还超过 10 则数组长度增加
	numLen := len(digits)

	// 处理空数组的边界情况
	if numLen == 0 {
		return []int{1}
	}

	for i := numLen - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] = 0
			continue
		} else {
			break
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 要求不使用额外数组空间，前提是有序的，则可以遍历时候，和前一位比较，如果重复，挪到后面，后续切掉
	// 双指针

	numLen := len(nums)
	if numLen == 0 {
		return 0
	}

	i := 0
	for j := 1; j < numLen; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 遍历数组，比较当前区间和前一个区间是否重叠，重叠判断依据为 前一个 end >= 后一个 start
	// 重叠则合并，不重叠则直接添加到结果数组

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])

	for _, interval := range intervals[1:] {
		// 结果中的最后一个区间
		last := &result[len(result)-1]

		// 结果中的区间 end 大于等于 当前区间 start 则重叠
		if (*last)[1] >= interval[0] {
			if (*last)[1] < interval[1] {
				(*last)[1] = interval[1]
			}
		} else {
			result = append(result, interval)
		}
	}

	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 首先剔除数组内大于 target 的整数，可以在初次遍历时直接过滤
	// 看测试 case 需要返回的是下标，题目中没说明，所以需要记录下标
	// 用 map 记录

	numMap := make(map[int]int)

	for i, num := range nums {
		if num > target {
			continue
		}

		diff := target - num

		if index, exists := numMap[diff]; exists {
			// 因为小下标最开始一定不在map中，所以此处的下标一定小于map中的下标
			return []int{index, i}
		}

		numMap[num] = i
	}

	return nil
}
