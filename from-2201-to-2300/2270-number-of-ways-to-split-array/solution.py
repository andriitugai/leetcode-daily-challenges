class Solution:
    def waysToSplitArray(self, nums: List[int]) -> int:
        leftSum, rightSum = 0, sum(nums)
        result = 0

        for num in nums[:-1]:
            leftSum += num
            rightSum -= num
            if leftSum >= rightSum:
                result += 1

        return result