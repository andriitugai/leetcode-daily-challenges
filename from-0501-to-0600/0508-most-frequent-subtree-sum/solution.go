/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findFrequentTreeSum(root *TreeNode) []int {
    sumCounts := make(map[int]int)

    var subtreeSum func(node *TreeNode) int
    subtreeSum = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        subSum := node.Val + subtreeSum(node.Left) + subtreeSum(node.Right)
        sumCounts[subSum] += 1
        return subSum
    }

    subtreeSum(root)
    maxCount := 0
    result := []int{}

    for ssum, count := range sumCounts {
        if count > maxCount {
            maxCount = count
            result = []int{ssum}
        } else if count == maxCount {
            result = append(result, ssum)
        }
    }

    return result
}