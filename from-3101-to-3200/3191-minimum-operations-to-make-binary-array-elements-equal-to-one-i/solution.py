class Solution:
    def minOperations(self, nums: List[int]) -> int:
        result = 0
        for i in range(len(nums) - 2):
            if nums[i] == 0:
                nums[i + 2] ^= 1
                nums[i + 1] ^= 1
                result += 1

        if nums[-2] == 0 or nums[-1] == 0:
            return -1

        return result