/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    var lca *TreeNode
    maxDepth := 0
    var dfs func(node *TreeNode, depth int) int
    dfs = func(node *TreeNode, depth int) int {
        if depth > maxDepth {
            maxDepth = depth
        }
        if node == nil {
            return depth
        }
        leftDepth := dfs(node.Left, depth + 1)
        rightDepth := dfs(node.Right, depth + 1)

        if leftDepth == rightDepth && leftDepth == maxDepth {
            lca = node
        }
        if leftDepth > rightDepth {
            return leftDepth
        }
        return rightDepth
    }
    dfs(root, 0)
    return lca
}