package main

import "fmt"

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

给定 target = 5，返回 true。

给定 target = 20，返回 false。

限制：
0 <= n <= 1000
0 <= m <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
 */
func main() {
	fmt.Println(findNumberIn2DArray([][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}},5))
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var i, j int = 0, len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		} else {
			return true
		}

	}
	return false
}