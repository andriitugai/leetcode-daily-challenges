class Solution:
    def minOperations(self, nums: List[int]) -> int:
        flipped = False
        result = 0
        for num in nums:
            if num == 0 and not flipped or num == 1 and flipped:
                result += 1
                flipped = not flipped
        return result