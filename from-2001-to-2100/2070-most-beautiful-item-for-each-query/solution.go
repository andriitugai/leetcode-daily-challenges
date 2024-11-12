func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] < items[j][0]
    })
    n := len(items)
    for i := 1; i < n; i ++ {
        if items[i][1] < items[i - 1][1] {
            items[i][1] = items[i - 1][1]
        }
    }

    result := make([]int, len(queries))
    for i, q := range(queries) {
        l, r := 0, n - 1
        for l <= r {
            m := (l + r) >> 1
            if q < items[m][0] {
                r = m - 1
            } else {
                l = m + 1
            }
        }
        if r >= 0 {
            result[i] = items[r][1]
        }
    }
    return result
}