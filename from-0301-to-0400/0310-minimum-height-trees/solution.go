func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }

    adj := make(map[int][]int)
    degree := make([]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
        degree[e[0]] += 1
        degree[e[1]] += 1
    }

    leaves := []int{}
    for i, d := range degree {
        if d == 1 {
            leaves = append(leaves, i)
        }
    }

    nodesLeft := n
    for nodesLeft > 2 {
        nL := len(leaves)
        nodesLeft -= nL
        for i := 0; i < nL; i ++ {
            leaf := leaves[0]
            leaves = leaves[1:]
            for _, nei := range adj[leaf] {
                degree[nei] -= 1
                if degree[nei] == 1 {
                    leaves = append(leaves, nei)
                }
            }
        }
    }
    return leaves
}