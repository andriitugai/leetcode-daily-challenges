/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sumEvenGrandparent(root *TreeNode) int {
    result := 0
    var dfs func(node *TreeNode, parent int, gparent int)
    dfs = func(node *TreeNode, parent int, gparent int){
        if node == nil {
            return
        }
        if gparent != -1 && gparent % 2 == 0 {
            result += node.Val
        }
        dfs(node.Left, node.Val, parent)
        dfs(node.Right, node.Val, parent)
    }
    dfs(root, -1, -1)
    return result
}