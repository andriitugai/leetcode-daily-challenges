func matrixSumQueries(n int, queries [][]int) int64 {
    rows, cols := n, n
    rSeen, cSeen := make(map[int]bool), make(map[int]bool)
    var result int64 = 0

    for i := len(queries) - 1; i >= 0; i -- {
        tpe, idx, val := queries[i][0], queries[i][1], queries[i][2]
        if tpe == 0 && !rSeen[idx] {
            result += int64(cols * val)
            rSeen[idx] = true
            rows -= 1
            // if rows == 0 {
            //     break
            // }
        } else if tpe == 1 && !cSeen[idx] {
            result += int64(rows * val)
            cSeen[idx] = true
            cols -= 1
            // if cols == 0 {
            //     break
            // }
        }
    }
    return result
}