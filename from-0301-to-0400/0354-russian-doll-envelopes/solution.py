class Solution:
    def maxEnvelopes(self, envelopes: List[List[int]]) -> int:
        envs = sorted(envelopes, key = lambda x: x[0])
        dp = [1] * len(envs)
        
        for i in range(1, len(envs)):
            for j in range(i):
                if envs[i][0] > envs[j][0] and envs[i][1] > envs[j][1] and dp[i] <= dp[j]:
                    dp[i] = dp[j] + 1
                                           
        return max(dp)