class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        mark = len(nums)
        for num in nums:
            nums[num%mark-1] += mark

        result = []
        for i in range(len(nums)):
            if nums[i] > 2 * mark:
                result.append(i+1)

        return result