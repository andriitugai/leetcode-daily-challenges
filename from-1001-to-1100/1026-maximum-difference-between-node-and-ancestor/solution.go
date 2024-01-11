/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxAncestorDiff(root *TreeNode) int {
    var dfs func(node *TreeNode, maxVal int, minVal int) int
    dfs = func(node *TreeNode, maxVal int, minVal int) int {
        if node == nil {
            return maxVal - minVal
        }
        if node.Val > maxVal {
            maxVal = node.Val
        }
        if node.Val < minVal {
            minVal = node.Val
        }
        lDiff := dfs(node.Left, maxVal, minVal)
        rDiff := dfs(node.Right, maxVal, minVal)
        if lDiff > rDiff {
            return lDiff
        }
        return rDiff
    }
    return dfs(root, root.Val, root.Val)
}