/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorder(root *TreeNode, str *strings.Builder){
	if root == nil{
		str.WriteString("#")
		return
	}

	str.WriteString(strconv.Itoa(root.Val))
	str.WriteString(",")
	
	preorder(root.Left, str)
	preorder(root.Right, str)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var p1 strings.Builder
	var p2 strings.Builder
   
    preorder(root, &p1)
	preorder(subRoot, &p2)

	fmt.Println(p1)
	fmt.Println(p2)
	
	return strings.Contains(p1.String(), p2.String())
}
