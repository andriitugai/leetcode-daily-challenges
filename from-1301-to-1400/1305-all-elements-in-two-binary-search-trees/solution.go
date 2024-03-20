/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    bst1 := []int{}
    bst2 := []int{}
    var dfs func(node *TreeNode, vals *[]int)
    dfs = func(node *TreeNode, vals *[]int) {
        if node == nil {
            return
        }
        dfs(node.Left, vals)
        *vals = append(*vals, node.Val)
        dfs(node.Right, vals)
        return
    }
    dfs(root1, &bst1)
    dfs(root2, &bst2)

    result := []int{}
    p1, p2 := 0, 0
    for p1 < len(bst1) && p2 < len(bst2) {
        if bst1[p1] < bst2[p2] {
            result = append(result, bst1[p1])
            p1 += 1
        } else {
            result = append(result, bst2[p2])
            p2 += 1
        }
    }
    result = append(result, bst1[p1:]...)
    result = append(result, bst2[p2:]...)
    return result
}