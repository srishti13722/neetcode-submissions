/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode){
	if root == nil{
		return
	}
    
	left := root.Left
	right := root.Right

	root.Left = right
	root.Right = left

	invertTree(root.Left)
	invertTree(root.Right)
}
func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}
