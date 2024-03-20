/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func balanceBST(root *TreeNode) *TreeNode {
    vals := []int{}
    var getVals func(node *TreeNode)
    getVals = func(node *TreeNode) {
        if node == nil {
            return
        }
        vals = append(vals, node.Val)
        getVals(node.Left)
        getVals(node.Right)
    }
    getVals(root)
    sort.Ints(vals)
    fmt.Println(vals)

    var buildBST func(nums []int) *TreeNode 
    buildBST = func(nums []int) *TreeNode {
        if len(nums) == 0 {
            return nil
        }
        mid := len(nums) / 2
        return &TreeNode{
            Val: nums[mid],
            Left: buildBST(nums[:mid]),
            Right: buildBST(nums[mid+1:]),
        }
    }
    return buildBST(vals)
}