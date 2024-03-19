/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    var insertMaxTree func(node *TreeNode, val int) *TreeNode
    insertMaxTree = func(node *TreeNode, val int) *TreeNode {
        if node == nil {
            return &TreeNode{Val: val}
        }
        if node.Val < val {
            return &TreeNode{Val: val, Left: node}
        }
        node.Right = insertMaxTree(node.Right, val)
        return node
    }
    return insertMaxTree(root, val)
}