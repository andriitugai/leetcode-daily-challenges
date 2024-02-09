class Solution(object):
    def largestDivisibleSubset(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        if not nums:
            return nums
        
        numss = sorted(nums)
        result = [[num] for num in numss]
        
        for i in range(len(numss)):
            for j in range(i):
                if numss[i] % numss[j] == 0 and len(result[i]) < len(result[j]) + 1:
                    result[i] = result[j] + [numss[i]]
                    
        return max(result, key=lambda x: len(x))