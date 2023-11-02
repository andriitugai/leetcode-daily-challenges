/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type treeInfo struct {
    sumTotal int
    numNodes int
}

func averageOfSubtree(root *TreeNode) int {
    var dfs func(node *TreeNode) treeInfo
    result := 0

    dfs = func(node *TreeNode) treeInfo {
        if node == nil {
            return treeInfo{0, 0}
        }
        leftTreeInfo := dfs(node.Left)
        rightTreeInfo := dfs(node.Right)

        numNodes := leftTreeInfo.numNodes + rightTreeInfo.numNodes + 1
        sumTotal := leftTreeInfo.sumTotal + rightTreeInfo.sumTotal + node.Val

        if numNodes > 0 && sumTotal / numNodes == node.Val {
            result += 1
        }

        return treeInfo{sumTotal:sumTotal, numNodes:numNodes}
    }

    dfs(root)
    return result
}