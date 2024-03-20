/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func distributeCoins(root *TreeNode) int {
    result := 0
    abs := func(a int) int {
        if a < 0 {
            return -a
        }
        return a
    }
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int{
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        result += abs(left) + abs(right)
        return node.Val + left + right - 1
    }
    dfs(root)
    return result
}