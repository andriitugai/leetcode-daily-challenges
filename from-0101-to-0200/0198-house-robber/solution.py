class Solution:
    def rob(self, nums: List[int]) -> int:
        dp = [0] * len(nums)
        
        for idx in range(len(nums)):
            to_rob = (dp[idx-2] + nums[idx]) if idx > 1 else nums[idx]
            not_to_rob = dp[idx-1] if idx > 0 else 0
            
            dp[idx] = max(to_rob, not_to_rob)
            
            
        return dp[-1]