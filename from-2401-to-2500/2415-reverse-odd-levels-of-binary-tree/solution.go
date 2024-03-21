/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func reverseOddLevels(root *TreeNode) *TreeNode {
    q := []*TreeNode{root}
    level := 0

    for len(q) > 0 {
        n := len(q)
        if level % 2 != 0 {
            left, right := 0, n - 1
            for left < right {
                q[left].Val, q[right].Val = q[right].Val, q[left].Val
                left += 1
                right -= 1
            }
        }
        for i := 0; i < n; i ++ {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        level += 1
    }
    return root
}