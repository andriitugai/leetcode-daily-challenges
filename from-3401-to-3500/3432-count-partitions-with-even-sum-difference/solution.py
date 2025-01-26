class Solution:
    def countPartitions(self, nums: List[int]) -> int:
        total = sum(nums)
        if total % 2:
            return 0

        curr = 0
        result = 0
        for num in nums[:-1]:
            curr += num
            total -= num
            if curr % 2 == total % 2:
                result += 1

        return result