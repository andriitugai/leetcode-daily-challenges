class Solution:
    def maxSum(self, nums: List[int]) -> int:
        mx = max(nums)
        if mx < 0:
            return mx
        return sum(num for num in set(nums) if num > 0)