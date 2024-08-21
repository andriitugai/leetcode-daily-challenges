class Solution:
    def strangePrinter(self, s: str) -> int:
        N = len(s)

        dp = [[sys.maxsize] * N for _ in range(N)]
        for i in range(N):
            dp[i][i] = 1

        for l in range(2, N+1):
            for i in range(N - l + 1):
                j = i + l - 1
                dp[i][j] = dp[i+1][j] + 1

                for k in range(i+1, j+1):
                    if s[i] == s[k]:
                        dp[i][j] = min(dp[i][j], dp[i][k-1] + (dp[k+1][j] if j > k else 0))

        return dp[0][N-1]