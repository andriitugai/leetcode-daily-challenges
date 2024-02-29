/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isEvenOddTree(root *TreeNode) bool {
    var n, prev, val int
    var node *TreeNode
    level := 0
    q := []*TreeNode{root}

    for len(q) > 0 {
        n = len(q)
        prev = q[0].Val
        for i := 0; i < n; i ++ {
            node = q[0]
            q = q[1:]
            val = node.Val
            if level % 2 == 0 && val % 2 == 0 || level % 2 != 0 && val % 2 != 0 {
                return false
            }
            if i > 0 {
                if level % 2 == 0 && val <= prev || level % 2 != 0 && val >= prev{
                    return false
                }
                prev = val
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        level += 1
    }
    return true
}