class Solution:
    def maximumTripletValue(self, nums: List[int]) -> int:
        maxi = float('-inf')
        diff = 0
        result = 0
        
        for i in range(len(nums)):
            maxi = max(maxi, nums[i])
            if i >= 2:
                result = max(result, diff * nums[i])
            if i >= 1:
                diff = max(diff, maxi - nums[i])
                
        return result