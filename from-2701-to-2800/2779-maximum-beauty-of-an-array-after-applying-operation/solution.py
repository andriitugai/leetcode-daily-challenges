class Solution:
    def maximumBeauty(self, nums: List[int], k: int) -> int:
        nums.sort()
        result = 0
        left = 0
        for right in range(len(nums)):
            while nums[right] - nums[left] > 2 * k:
                left += 1

            result = max(result, right - left + 1)

        return result