class Solution:
    def numRollsToTarget(self, n: int, k: int, target: int) -> int:
        dp = {}
        def ways(dice, target):
            nonlocal dp, k
            if dice == 0:
                return 1 if target == 0 else 0
            if (dice, target) in dp:
                return dp[(dice, target)]
            
            # dp(d, f, target) = dp(d-1, f, target-1) + dp(d-1, f, target-2) + ... + dp(d-1, f, target-f)
            res = 0
            for val in range(1, k+1):
                res += ways(dice - 1, target - val)
                
            dp[(dice, target)] = res
            return res
        
        return ways(n, target) % (1_000_000_000 + 7)
    