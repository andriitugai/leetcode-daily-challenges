func minFallingPathSum(grid [][]int) int {
    dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid))
	}
	copy(dp[0], grid[0])
	for i := 1; i < len(grid); i++ {
		prev := make([]int, len(grid))
		copy(prev, dp[i-1])
		sort.Ints(prev)
		for j := 0; j < len(grid); j++ {
			if dp[i-1][j] == prev[0] {
				dp[i][j] = grid[i][j] + prev[1]
			} else {
				dp[i][j] = grid[i][j] + prev[0]
			}
		}
	}
	return min(dp[len(grid)-1]...)
}

func min(values ...int) int {
	minValue := math.MaxInt64
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}