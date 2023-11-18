func findColumnWidth(grid [][]int) []int {
    width := func(a int) int {
        if a == 0 {
            return 1
        }
        result := 0
        if a < 0 {
            result += 1
        }

        for a != 0 {
            a /= 10
            result += 1
        }
        return result
    }
    result := []int{}
    for c := 0; c < len(grid[0]); c ++ {
        maxW := 0
        for r := 0; r < len(grid); r ++ {
            w := width(grid[r][c])
            if w > maxW {
                maxW = w
            }
        }
        result = append(result, maxW)
    }
    return result
}