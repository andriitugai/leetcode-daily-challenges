/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func replaceValueInTree(root *TreeNode) *TreeNode {
    q := []*TreeNode{root}
    var n int
    level := 0
    levelSums := make(map[int]int)

    for len(q) > 0 {
        n = len(q)
        for i := 0; i < n; i ++ {
            node := q[0]
            q = q[1:]
            levelSums[level] += node.Val
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        level += 1
    }

    var dfs func(node *TreeNode, lvl int, siblingSum int)
    dfs = func(node *TreeNode, lvl int, siblingSum int) {
        if node == nil {
            return
        }
        node.Val = levelSums[lvl] - siblingSum
        ss := 0
        if node.Left != nil {
            ss += node.Left.Val
        }
        if node.Right != nil {
            ss += node.Right.Val
        }
        dfs(node.Left, lvl + 1, ss)
        dfs(node.Right, lvl + 1, ss)
    }
    dfs(root, 0, root.Val)
    return root
}