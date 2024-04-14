/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumOfLeftLeaves(root *TreeNode) int {
    var dfs func(node *TreeNode, isLeft bool) int
    dfs = func(node *TreeNode, isLeft bool) int {
        if isLeft && node.Left == nil && node.Right == nil {
            return node.Val
        }
        result := 0
        if node.Left != nil {
            result += dfs(node.Left, true)
        }
        if node.Right != nil {
            result += dfs(node.Right, false)
        }
        return result
    }
    return dfs(root, false)
}