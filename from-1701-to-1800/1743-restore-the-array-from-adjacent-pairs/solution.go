func restoreArray(adjacentPairs [][]int) []int {
    adj := make(map[int][]int)
    for _, ap := range adjacentPairs {
        adj[ap[0]] = append(adj[ap[0]], ap[1])
        adj[ap[1]] = append(adj[ap[1]], ap[0])
    }
    res := make([]int, len(adjacentPairs) + 1)
    // Look for starting elem
    for el, adjList := range adj {
        if len(adjList) == 1 {
            res[0] = el
            res[1] = adjList[0]
            break
        }
    }
    for i := 1; i < len(res) - 1; i++ {
        for _, nbr := range adj[res[i]] {
            if nbr != res[i - 1] {
                res[i + 1] = nbr
                break
            }
        }
    }
    return res
}