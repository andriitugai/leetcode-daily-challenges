class Solution:
    def hasIncreasingSubarrays(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        curLen = 1
        prevLen = 0
        maxGrowLen = 0
        for i in range(1, n):
            if nums[i] > nums[i - 1]:
                curLen += 1
            else:
                prevLen = curLen
                curLen = 1
            
            maxGrowLen = max(maxGrowLen, curLen // 2, min(curLen, prevLen))
        return maxGrowLen >= k