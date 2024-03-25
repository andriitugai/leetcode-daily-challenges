/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    type result struct {
        node *TreeNode
        dist int
    }
    var dfs func(node *TreeNode) result
    dfs = func(node *TreeNode) result {
        if node == nil {
            return result{node: nil, dist: 0}
        }
        left, right := dfs(node.Left), dfs(node.Right)
        if left.dist > right.dist {
            return result{node: left.node, dist: left.dist + 1}
        }
        if right.dist > left.dist {
            return result{node: right.node, dist: right.dist + 1}
        }
        return result{node: node, dist: left.dist + 1}
    }
    return dfs(root).node
}