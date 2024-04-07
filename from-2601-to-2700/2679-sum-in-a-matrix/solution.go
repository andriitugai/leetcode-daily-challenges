func matrixSum(nums [][]int) int {
    m, n := len(nums), len(nums[0])
    for k := 0; k < m; k ++ {
        sort.Slice(nums[k], func(i, j int) bool {
            return nums[k][i] > nums[k][j]
        })
    }
    total := 0
    for col := 0; col < n; col ++ {
        score := 0
        for row := 0; row < m; row ++ {
            if nums[row][col] > score {
                score = nums[row][col]
            }
        }
        total += score
    }
    return total
}