class Solution:
    def sumOfGoodNumbers(self, nums: List[int], k: int) -> int:
        n = len(nums)
        result = 0
        for i, num in enumerate(nums):
            if i - k >= 0 and nums[i - k] >= num:
                continue
            if i + k < n and nums[i + k] >= num:
                continue
            result += num

        return result