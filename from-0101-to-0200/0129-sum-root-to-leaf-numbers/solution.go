/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumNumbers(root *TreeNode) int {
    result := 0
    var dfs func(node *TreeNode, curr int)
    dfs = func(node *TreeNode, curr int) {
        if node == nil {
            return
        }
        curr = curr * 10 + node.Val
        if node.Left == nil && node.Right == nil {
            result += curr
        }
        dfs(node.Left, curr)
        dfs(node.Right, curr)
    }
    dfs(root, 0)
    return result
}