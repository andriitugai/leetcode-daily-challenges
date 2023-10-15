class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        m = len(grid)
        n = len(grid[0])

        dp = [ [0] * n for _ in range(m) ]
        t = 0
        for row in range(m):
            t += grid[row][0]
            dp[row][0] = t
        
        t = 0
        for col in range(n):
            t += grid[0][col]
            dp[0][col] = t

        for r in range(1, m):
            for c in range(1, n):
                dp[r][c] = min(dp[r-1][c], dp[r][c-1]) + grid[r][c]

        return dp[-1][-1]