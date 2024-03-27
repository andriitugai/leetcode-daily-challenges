class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        if k < 1:
            return 0

        result = 0
        product = 1
        left, right = 0, 0

        while right < len(nums):
            product *= nums[right]
            
            while product >= k and left <= right:
                product //= nums[left]
                left += 1

            result += (right - left + 1)
            right += 1

        return result