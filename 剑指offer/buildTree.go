package main
/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]

返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
 */
func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preLen int = len(preorder)
	if preLen == 0 {
		return  nil
	}
	rootIndex := 0
	for index,value := range inorder{
		if value == preorder[0]{
			rootIndex = index
			break
		}
	}
	root := &TreeNode{Val:preorder[0]}
	root.Left = buildTree(preorder[1:rootIndex+1],inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:],inorder[rootIndex+1:])
	return root
}