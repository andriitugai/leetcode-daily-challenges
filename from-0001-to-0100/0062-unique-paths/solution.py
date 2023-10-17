class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [ [1 if _c == 0 or _r == 0 else 0 for _c in range(n)] for _r in range(m) ]
        
        for r in range(1, m):
            for c in range(1, n):
                dp[r][c] = dp[r-1][c] + dp[r][c-1]
                
        return dp[m-1][n-1]