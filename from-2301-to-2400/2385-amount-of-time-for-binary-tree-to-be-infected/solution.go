/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func amountOfTime(root *TreeNode, start int) int {
    g := make(map[int][]int)
    var treeToGraph func(node *TreeNode)
    treeToGraph = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left != nil {
            g[node.Val] = append(g[node.Val], node.Left.Val)
            g[node.Left.Val] = append(g[node.Left.Val], node.Val)
            treeToGraph(node.Left)
        }
        if node.Right != nil {
            g[node.Val] = append(g[node.Val], node.Right.Val)
            g[node.Right.Val] = append(g[node.Right.Val], node.Val)
            treeToGraph(node.Right)
        }
    }
    treeToGraph(root)

    result := 0
    q := []int{}
    visited := make(map[int]bool)
    visited[start] = true
    for _, n := range g[start] {
        q = append(q, n)
    }

    for len(q) > 0 {
        l := len(q)
        for i := 0; i < l; i ++ {
            node := q[0]
            q = q[1:]
            visited[node] = true
            for _, n := range g[node] {
                if !visited[n] {
                    q = append(q, n)
                }   
            }
        }
        result += 1
    }
    return result
}