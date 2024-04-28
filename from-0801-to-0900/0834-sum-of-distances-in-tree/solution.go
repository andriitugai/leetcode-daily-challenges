func sumOfDistancesInTree(n int, edges [][]int) []int {
    res := make([]int, n)
    count := make([]int, n)
    tree := make([][]int, n)
    for _, edge := range edges {
        tree[edge[0]] = append(tree[edge[0]], edge[1])
        tree[edge[1]] = append(tree[edge[1]], edge[0])
    }
    dfs(0, -1, tree, res, count)
    dfs2(0, -1, tree, res, count)
    return res
}


func dfs(root int, pre int, tree [][]int, res []int, count []int) {
    for _, c := range tree[root] {
        if c == pre {
            continue
        }
        dfs(c, root, tree, res, count)
        count[root] += count[c]
        res[root] += count[c] + res[c]
    }
    count[root]++
}

func dfs2(root int, pre int, tree [][]int, res []int, count []int) {
    for _, c := range tree[root] {
        if c == pre {
            continue
        }
        res[c] = res[root] - count[c] * 2 + len(count)
        dfs2(c, root, tree, res, count) 
    }
}