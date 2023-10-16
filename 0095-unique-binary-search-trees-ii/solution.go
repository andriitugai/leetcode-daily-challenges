/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func generateTrees(n int) []*TreeNode {
    return generate_(1, n)
}

func generate_(start int, end int) []*TreeNode {
    var nodes []*TreeNode

    if start > end {
        nodes = append(nodes, nil)
        return nodes
    }

    for val:= start; val <= end; val ++ {
        leftTrees := generate_(start, val - 1)
        rightTrees := generate_(val + 1, end)

        for _, left := range leftTrees {
            for _, right := range rightTrees {
                newNode := &TreeNode {
                    Val: val,
                    Left: left,
                    Right: right,
                }
                nodes = append(nodes, newNode)
            }
        }
    }

    return nodes
}