package main

import "fmt"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof

 */
func main() {
	testNode := &ListNode{1,nil}
	data1 := &ListNode{2,nil}
	data2 := &ListNode{3,nil}
	testNode.Next = data1
	data1.Next = data2
	fmt.Println(reversePrint(testNode))
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	headbak := head
	var i int = 0
	for head.Next != nil {
		head = head.Next
		i++
	}
	arr := make([]int, i+1)
	for i >= 0 {
		arr[i] = headbak.Val
		headbak = headbak.Next
		i--
	}
	return arr
}