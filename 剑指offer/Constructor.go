package main

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

示例 2：
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
 */
func main() {

}
type CQueue struct {
	in []int
	out []int
}
func Constructor() CQueue {
	return CQueue{}
}
func (this *CQueue) AppendTail(value int)  {
	this.in = append(this.in,value)
}

func (this *CQueue) DeleteHead() int {
	var outLen int = len(this.out)
	inLen := len(this.in)
	if inLen  == 0 && outLen == 0{
		return -1
	}
	if outLen == 0 {
		for i := inLen;i > 0; i--{
			var outIndex int = i - 1
			var outValue int = this.in[outIndex]
			this.out = append(this.out,outValue)
			this.in = this.in[:outIndex]
		}
	}
	var newOutLen int = len (this.out)
	var outValue int = this.out[newOutLen-1]
	this.out = this.out[:newOutLen-1]
	return outValue
}