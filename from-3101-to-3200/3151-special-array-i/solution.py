class Solution:
    def isArraySpecial(self, nums: List[int]) -> bool:
        prev = nums[0] & 1
        for num in nums[1:]:
            curr = num & 1
            if curr == prev:
                return False
            prev = curr
        return True