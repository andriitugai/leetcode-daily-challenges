class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:

        n = len(obstacleGrid)
        m = len(obstacleGrid[0])
        
        if obstacleGrid[0][0] == 1 or obstacleGrid[n-1][m-1] == 1:
            return 0
        
        dp = [ [0] * m for _ in range(n)]
        # dp[0][0] = 1
        
        for r in range(n):
            for c in range(m):
                if obstacleGrid[r][c]:
                    dp[r][c] = 0
                else:
                    if r == 0 and c == 0:
                        dp[r][c] = 1
                    elif r == 0 and c > 0:
                        dp[r][c] = dp[r][c-1]
                    elif c == 0 and r > 0:
                        dp[r][c] = dp[r-1][c]
                    else:
                        dp[r][c] = dp[r - 1][c] + dp[r][c - 1]
                        
        return dp[n - 1][m - 1]
 