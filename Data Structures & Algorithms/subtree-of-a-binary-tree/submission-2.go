/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorder(root *TreeNode)string{
	if root == nil{
		return "null"
	}

	var str strings.Builder
	str.WriteString(strconv.Itoa(root.Val))
	str.WriteString(preorder(root.Left))
	str.WriteString(preorder(root.Right))
	return str.String()
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    p1 := preorder(root)
	p2 := preorder(subRoot)

	fmt.Println(p1)
	fmt.Println(p2)
	
	return strings.Contains(p1, p2)
}
