class Solution:
    def __init__(self):
        self.dp = [0]
        self.squares = [i*i for i in range(1, 101)]

    def numSquares(self, n: int) -> int:
        if len(self.dp) > n:
            return self.dp[n]

        while len(self.dp) < n + 1:
            dpt = inf
            for s in self.squares:
                if s > len(self.dp):
                    break
                dpt = min(dpt, 1 + self.dp[-s])
            self.dp.append(dpt)

        return self.dp[n]