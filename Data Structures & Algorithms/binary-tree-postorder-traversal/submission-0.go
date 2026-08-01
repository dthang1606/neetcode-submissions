/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
    var dfs func(node *TreeNode)
	var result []int
	dfs = func(node *TreeNode){
		if node == nil{
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}
