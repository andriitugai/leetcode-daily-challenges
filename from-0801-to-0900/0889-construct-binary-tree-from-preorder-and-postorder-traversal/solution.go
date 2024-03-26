/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    var construct func(pre, post []int) *TreeNode
    construct = func(pre, post []int) *TreeNode{
        if len(pre) == 0 {
            return nil
        }
        root := &TreeNode{Val: pre[0]}
        if len(pre) == 1 {
            return root
        }
        leftRootVal := pre[1]
        leftSubtreeSize := slices.Index(post, leftRootVal) + 1
        root.Left = construct(pre[1:leftSubtreeSize+1], post[:leftSubtreeSize])
        root.Right = construct(pre[leftSubtreeSize+1:], post[leftSubtreeSize:])
        return root
    }
    return construct(preorder, postorder)
}