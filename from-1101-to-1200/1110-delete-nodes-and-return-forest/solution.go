/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    toDelete := make(map[int]bool)
    for _, val := range to_delete {
        toDelete[val] = true
    }
    result := []*TreeNode{}

    var dfs func(node *TreeNode, isRoot bool) *TreeNode
    dfs = func(node *TreeNode, isRoot bool) *TreeNode {
        if node == nil {
            return nil
        }
        nodeDeleted := toDelete[node.Val]
        if isRoot && !nodeDeleted {
            result = append(result, node)
        }
        node.Left = dfs(node.Left, nodeDeleted)
        node.Right = dfs(node.Right, nodeDeleted)
        if nodeDeleted {
            return nil
        }
        return node
    }
    dfs(root, true)
    return result
}