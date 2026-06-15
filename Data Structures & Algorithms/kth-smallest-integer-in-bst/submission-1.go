/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(root *TreeNode, ans *int, count *int, k int) bool{
	if root == nil{
		return false
	}

	if inorder(root.Left, ans, count, k){
		return true
	
	}
	*count++
	if *count == k{
		*ans = root.Val
		return true
	}
	return inorder(root.Right, ans, count, k)
}
func kthSmallest(root *TreeNode, k int) int {
	ans , count := 0 , 0
    inorder(root, &ans, &count, k)
	return ans
}
