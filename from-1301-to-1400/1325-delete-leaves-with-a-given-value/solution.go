/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfs func(node *TreeNode, tgt int) bool 
    dfs = func(node *TreeNode, tgt int) bool {
        if node == nil {
            return true
        }
        if dfs(node.Left, tgt) {
            node.Left = nil
        }
        if dfs(node.Right, tgt) {
            node.Right = nil
        }
        if node.Val == tgt && node.Left == nil && node.Right == nil {
            return true
        }
        return false
    }
    dfs(root, target)
    if root.Left == nil && root.Right == nil && root.Val == target {
        return nil
    }
    return root
}