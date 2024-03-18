/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructMaximumBinaryTree(nums []int) *TreeNode {
    var maxBTree func(arr []int) *TreeNode
    maxBTree = func(arr []int) *TreeNode {
        if len(arr) == 0 {
            return nil
        }
        maxVal := -1
        maxIdx := -1
        for i := 0; i < len(arr); i ++ {
            if arr[i] > maxVal {
                maxVal = arr[i]
                maxIdx = i
            }
        }
        return &TreeNode {
            Val: maxVal,
            Left: maxBTree(arr[:maxIdx]),
            Right: maxBTree(arr[maxIdx+1:]),
        }
    }
    return maxBTree(nums)
}