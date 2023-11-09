**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node.Val == 0 {
            return false
        } else if node.Val == 1 {
            return true
        } else if node.Val == 2 {
            return dfs(node.Left) || dfs(node.Right)
        } else {
            return dfs(node.Left) && dfs(node.Right)
        }
    }
    return dfs(root)
}