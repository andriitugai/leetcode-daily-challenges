/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func largestValues(root *TreeNode) []int {
    result := []int{}
    q := []*TreeNode{}
    if root != nil {
        q = append(q, root)
    }

    for len(q) > 0 {
        lenRow := len(q)
        rowMax := math.MinInt
        for i := 0; i < lenRow; i++ {
            node := q[0]
            q = q[1:]
            if node.Val > rowMax {
                rowMax = node.Val
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        result = append(result, rowMax)
    }

    return result
}