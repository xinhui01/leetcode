package main

import (
	"fmt"
	"sort"
)
/**
最接近三数之和
来源：力扣（LeetCode）
链接:https://leetcode-cn.com/problems/3sum-closest
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
提示：
    3 <= nums.length <= 10^3
    -10^3 <= nums[i] <= 10^3
    -10^4 <= target <= 10^4
 */
func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4},0))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	tmpsum := nums[0] + nums[1] + nums[2]
	tmpmin := tmpsum - target
	if tmpmin == 0 {
		return tmpsum
	}
	if tmpmin < 0 {
		tmpmin = tmpmin * -1
	}
	for index, _ := range nums {
		left := 0
		right := len(nums) - 1
		if index > 0 && nums[index] == nums[index-1] {
			left = index - 1
		}
		for left < index && index < right {
			sum := nums[left] + nums[index] + nums[right]
			min := sum - target
			if min == 0 {
				return sum
			} else if min > 0 {
				right--
			} else {
				left++
				min = min * -1
			}
			if min < tmpmin {
				tmpmin = min
				tmpsum = sum
			}
		}
	}
	return tmpsum
}