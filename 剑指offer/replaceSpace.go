package main

import "fmt"

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."

限制：
0 <= s 的长度 <= 10000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
 */
func main() {
	fmt.Println(replaceSpace("We are happy."))
}
func replaceSpace(s string) string {
	var str string = ""
	for _, val := range s {
		if val == ' ' {
			str += "%20"
		} else {
			str += string(val)
		}
	}
	return str
}