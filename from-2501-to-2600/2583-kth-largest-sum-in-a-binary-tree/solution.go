/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthLargestLevelSum(root *TreeNode, k int) int64 {
    q := []*TreeNode{root}
    sums := []int64{}
    var lvlSum int64

    for len(q) > 0 {
        n := len(q)
        lvlSum = 0
        for i := 0; i < n; i ++ {
            node := q[0]
            q = q[1:]
            lvlSum += int64(node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        sums = append(sums, lvlSum)
    }
    if len(sums) < k {
        return -1
    }
    sort.Slice(sums, func(i, j int) bool{
        return sums[i] > sums[j]
    })
    return sums[k - 1]
}