class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        additions = {}
        for i,num in enumerate(nums):
            if target-num in additions:
                return [additions[target-num], i]
            else:
                additions[num] = i