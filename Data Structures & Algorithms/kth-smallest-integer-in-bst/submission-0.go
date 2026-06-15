/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(root *TreeNode, ans *int, count *int, k int){
	if root == nil{
		return
	}

	inorder(root.Left, ans, count, k)
	*count++
	if *count == k{
		*ans = root.Val
		return
	}
	inorder(root.Right, ans, count, k)
}
func kthSmallest(root *TreeNode, k int) int {
	ans , count := 0 , 0
    inorder(root, &ans, &count, k)
	return ans
}
