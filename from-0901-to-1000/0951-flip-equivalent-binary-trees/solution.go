/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    var isFlippable func(node1, node2 *TreeNode) bool
    isFlippable = func(node1, node2 *TreeNode) bool {
        if node1 == nil || node2 == nil {
            return node1 == node2
        }
        return node1.Val == node2.Val &&
               (isFlippable(node1.Left, node2.Left) && isFlippable(node1.Right, node2.Right) ||
               isFlippable(node1.Left, node2.Right) && isFlippable(node1.Right, node2.Left))
    }
    return isFlippable(root1, root2)
}