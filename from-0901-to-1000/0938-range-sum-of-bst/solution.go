/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rangeSumBST(root *TreeNode, low int, high int) int {
    result := 0
    var dfs func(node TreeNode)
    dfs = func(node TreeNode) {
        if low <= node.Val && node.Val <= high {
            result += node.Val
        } 
        if node.Val > low && node.Left != nil {
            dfs(*node.Left)
        } 
        if node.Val < high && node.Right != nil {
            dfs(*node.Right)
        }
    }
    dfs(*root)
    return result
}