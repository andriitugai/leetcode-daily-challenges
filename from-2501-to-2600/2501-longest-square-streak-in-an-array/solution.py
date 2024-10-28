class Solution:
    def longestSquareStreak(self, nums: List[int]) -> int:
        exists = dict()
        for num in nums:
            exists[num] = True

        visited = set()
        maxLen = -1
        nums.sort()
        maxNum = nums[len(nums) - 1]
        for num in nums:
            if num in visited:
                continue
            
            square = num
            curLen = 0
            while square <= maxNum and exists.get(square, False):
                visited.add(square)
                curLen += 1
                square *= square
            if curLen > 1 and curLen > maxLen:
                maxLen = curLen

        return maxLen