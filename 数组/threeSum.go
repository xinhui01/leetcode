package main

import (
	"fmt"
	"sort"
)
/**
三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */
func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println(nums)
	var saveArr [][]int
	lens := len(nums) - 1
	for index, _ := range nums {
		left := 0
		right := lens
		//避免重复计算
		if index > 0 && nums[index] == nums[index-1] {
			left = index - 1
		}
		for left < index && index < right {
			sum := nums[left] + nums[right] + nums[index]
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < lens && nums[right] == nums[right+1] {
				right--
				continue
			}
			if sum == 0 {
				saveArr = append(saveArr, []int{nums[index], nums[left], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}

		}
	}
	return saveArr
}