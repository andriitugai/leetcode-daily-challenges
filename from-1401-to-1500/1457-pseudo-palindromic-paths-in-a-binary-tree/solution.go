/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pseudoPalindromicPaths (root *TreeNode) int {
    result := 0
    var dfs func(node *TreeNode, seen map[int]bool)
    dfs  = func(node *TreeNode, seen map[int]bool){
        if seen[node.Val] {
            delete(seen, node.Val)
        } else {
            seen[node.Val] = true
        }

        seenCopy := map[int]bool{}
        for k, v := range seen {
            seenCopy[k] = v
        }
        if node.Left != nil {
            dfs(node.Left, seenCopy)
        }

        seenCopy = map[int]bool{}
        for k, v := range seen {
            seenCopy[k] = v
        }
        if node.Right != nil {
            dfs(node.Right, seenCopy)
        }

        if node.Left == nil && node.Right == nil {
            if len(seen) < 2 {
                result += 1
            }
        }
    }
    dfs(root, map[int]bool{})
    return result
}