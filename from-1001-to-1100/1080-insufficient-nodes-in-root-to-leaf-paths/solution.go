/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    var dfs func(node *TreeNode, curSum int) bool
    dfs = func(node *TreeNode, curSum int) bool {
        curSum += node.Val

        if node.Left == nil && node.Right == nil {
            return curSum >= limit
        }
        left := false
        if node.Left != nil {
            left = dfs(node.Left, curSum)
        }
        right := false
        if node.Right != nil {
            right = dfs(node.Right, curSum)
        }
        if !left {
            node.Left = nil
        }
        if !right {
            node.Right = nil
        }
        return left || right
    }
    if dfs(root, 0) {
        return root
    }
    return nil
}