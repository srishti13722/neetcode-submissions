/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    var queue []*TreeNode

	if root != nil{
		queue = append(queue, root)
	}

	var res [][]int

	for len(queue) > 0{
		size := len(queue)
		var temp []int

		for i := 0 ; i < size ; i++{
			top := queue[0]
			queue = queue[1:]
			temp = append(temp, top.Val)

			if top.Left != nil{
				queue = append(queue, top.Left)
			}

			if top.Right != nil{
				queue = append(queue, top.Right)
			}
		}

		res = append(res, temp)
	}

	return res
}
