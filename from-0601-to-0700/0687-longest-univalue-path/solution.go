/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func longestUnivaluePath(root *TreeNode) int {
    maxL := 0
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    var longestPath func(node *TreeNode) int
    longestPath = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := longestPath(node.Left)
        right := longestPath(node.Right)
        leftSide, rightSide := 0, 0
        if node.Left != nil && node.Left.Val == node.Val {
            leftSide += 1 + left
        }
        if node.Right != nil && node.Right.Val == node.Val {
            rightSide += 1 + right
        }
        maxL = max(maxL, leftSide + rightSide)
        return max(leftSide, rightSide)
    }
    longestPath(root)
    return maxL
}