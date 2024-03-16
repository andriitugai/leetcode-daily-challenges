class Solution(object):
    def findMaxLength(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        
        ans = 0
        # count 1s as +1 and 0s as -1
        counts = {1: 1, 0: -1}
        count = 0
        idx = 0
        
        results = { 0: -1}  # map curr count to the first index where it has been reached
                            # our count starts at -1 with the result 0
        while idx < len(nums):
            count += counts[nums[idx]]
            if count in results:
                ans = max(ans, idx - results[count])
            else:
                results[count] = idx
                
            idx += 1
            
        return ans
        
        