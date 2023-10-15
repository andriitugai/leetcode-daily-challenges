class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        cmax = 0
        gmax = float('-inf')

        for num in nums:
            cmax = max(num, cmax + num)
            gmax = max(cmax, gmax)

        return gmax