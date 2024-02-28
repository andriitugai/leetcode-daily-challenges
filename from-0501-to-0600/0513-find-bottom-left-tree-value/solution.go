/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findBottomLeftValue(root *TreeNode) int {
    q := []*TreeNode{root}
    
    var leftmost, n int
    var node *TreeNode

    for len(q) > 0 {
        leftmost = q[0].Val
        n = len(q)
        for i := 0; i < n; i ++ {
            node = q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
    return leftmost
}