class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        
        fast = 1
        slow = 1
        prev = nums[0]
        count = 1
        
        while fast < len(nums):
            if nums[fast] != prev:
                prev = nums[fast]
                nums[slow] = prev
                slow += 1
                count = 1
            else:
                count += 1
                if count <= 2:
                    nums[slow] = prev
                    slow += 1
                    count += 1                    
                
            fast += 1
            
        return slow