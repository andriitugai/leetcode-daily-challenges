/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        return &TreeNode{Val: val, Left: root}
    }
    var node *TreeNode
    q := []*TreeNode{root}
    lvl := 1
    for lvl < depth - 1 {
        size := len(q)
        for i := 0; i < size; i ++ {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        lvl += 1
    }

    for len(q) > 0 {
        node = q[0]
        q = q[1:]
        node.Left = &TreeNode{Val: val, Left: node.Left}
        node.Right = &TreeNode{Val: val, Right: node.Right}
    }
    return root
}