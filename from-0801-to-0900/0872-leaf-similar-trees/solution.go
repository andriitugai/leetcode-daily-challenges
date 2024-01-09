/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var getLeafVals func(node *TreeNode, leafVals *[]int)
    getLeafVals = func(node *TreeNode, leafVals *[]int) {
        if node == nil {
            return
        }
        if node.Left == nil && node.Right == nil {
            *leafVals = append(*leafVals, node.Val)
            return
        } 
        getLeafVals(node.Left, leafVals)
        getLeafVals(node.Right, leafVals)
    }

    leafVals1, leafVals2 := []int{}, []int{}
    getLeafVals(root1, &leafVals1)
    getLeafVals(root2, &leafVals2)

    return reflect.DeepEqual(leafVals1, leafVals2)
}