func canCross(stones []int) bool {
	dp := make(map[int]map[int]int, len(stones))
	for _, v := range stones {
		dp[v] = make(map[int]int)
	}
	dp[0][0] = 0
	for i := 0; i < len(stones); i++ {
		for _, k := range dp[stones[i]] {
			for step := k - 1; step <= k+1; step++ {
				if step > 0 {
					if _, ok := dp[stones[i]+step]; ok {
						dp[stones[i]+step][i] = step
					}
				}
			}
		}
	}

	return len(dp[stones[len(stones)-1]]) != 0
}