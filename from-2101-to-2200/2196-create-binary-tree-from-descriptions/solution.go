/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode)
    hasParent := make(map[int]bool)

    var node, cnode *TreeNode

    for _, d := range descriptions {
        parent, child, left := d[0], d[1], d[2]
        if _, ok := nodes[parent]; !ok {
            node = &TreeNode{Val: parent}
            nodes[parent] = node
        } else {
            node = nodes[parent]
        }

        if _, ok := nodes[child]; !ok {
            cnode = &TreeNode{Val: child}
            nodes[child] = cnode
        } else {
            cnode = nodes[child]
        }

        if left == 1 {
            node.Left = cnode
        } else {
            node.Right = cnode
        }
        hasParent[child] = true
    }

    for par, _ := range nodes {
        if !hasParent[par] {
            return nodes[par]
        }
    }
    return nil
}