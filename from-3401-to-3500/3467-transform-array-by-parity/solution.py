class Solution:
    def transformArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0] * n
        ptr = n - 1
        for num in nums:
            if num % 2:
                result[ptr] = 1
                ptr -= 1
        return result