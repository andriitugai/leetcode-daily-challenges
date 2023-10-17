func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    findRoot := func() int {
        children := make(map[int]bool)
        for _, lc := range leftChild {
            children[lc] = true
        }
        for _, rc := range rightChild {
            children[rc] = true
        }

        for i := 0; i < n; i++ {
            if !children[i] {
                return i
            }
        }
        return -1
    }
    root := findRoot()
    if root == -1 {
        return false
    }

    stack := []int{root}
    visited := make(map[int]bool)

    for len(stack) != 0 {
        node := stack[0]
        stack = stack[1:]

        if visited[node] {
            return false
        }
        visited[node] = true

        if leftChild[node] != -1 {
            stack = append(stack, leftChild[node])
        }
        if rightChild[node] != -1 {
            stack = append(stack, rightChild[node])
        }
    }
    return len(visited) == n
}