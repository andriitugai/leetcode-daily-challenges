class Solution:
    def maxAdjacentDistance(self, nums: List[int]) -> int:
        max_adj = 0
        n = len(nums)
        for i in range(n - 1):
            max_adj = max(max_adj, abs(nums[i] - nums[i+1]))

        return max(max_adj, abs(nums[0] - nums[n - 1]))