class Solution:
    def countVowelPermutation(self, n: int) -> int:
        a, e, i, o, u = 0, 1, 2, 3, 4
        dp_prev = [1, 1, 1, 1, 1]
        
        for _ in range(n-1):
            dp_curr = [0, 0, 0, 0, 0]
            dp_curr[a] = dp_prev[e] + dp_prev[i] + dp_prev[u]
            dp_curr[e] = dp_prev[i] + dp_prev[a]
            dp_curr[i] = dp_prev[e] + dp_prev[o]
            dp_curr[o] = dp_prev[i]
            dp_curr[u] = dp_prev[i] + dp_prev[o]
            
            dp_prev = dp_curr
            
        return sum(dp_prev) % (10**9 + 7)