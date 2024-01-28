func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m, n := len(matrix), len(matrix[0])

    for row := 0; row < m; row ++ {
        for col := 1; col < n; col ++ {
            matrix[row][col] += matrix[row][col - 1]
        }
    }

    result := 0
    for c1 := 0; c1 < n; c1 ++ {
        for c2 := c1; c2 < n; c2 ++ {
            mp := make(map[int]int)
            mp[0] = 1
            sum := 0

            for row := 0; row < m; row ++ {
                sum += matrix[row][c2]
                if c1 > 0 {
                    sum -= matrix[row][c1 - 1]
                }
                result += mp[sum - target]
                mp[sum] += 1
            }
        }
    }
    return result
}