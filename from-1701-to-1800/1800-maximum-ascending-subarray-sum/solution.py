class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        maxsum = 0
        cursum = 0
        prev = 0
        for curr in nums:
            if curr > prev:
                cursum += curr
            else:
                if cursum > maxsum:
                    maxsum = cursum
                cursum = curr
            prev = curr
            
        return max(maxsum, cursum)
        