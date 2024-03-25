/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func recoverTree(root *TreeNode)  {
    var dfs func(node *TreeNode)
    var n1, n2, parent *TreeNode = nil, nil, nil
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if parent != nil && node.Val < parent.Val {
            if n1 == nil {
                n1 = parent
            }
            n2 = node
        }
        parent = node
        dfs(node.Right)
    }
    dfs(root)
    n1.Val, n2.Val = n2.Val, n1.Val
}