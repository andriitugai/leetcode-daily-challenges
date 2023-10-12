class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if not nums:
            return 0
        
        last = len(nums)-1
        fast = 0
        while nums[last] == val and last > 0:
            last -= 1
            
        while fast <= last:
            if nums[fast] == val:
                nums[fast], nums[last] = nums[last], nums[fast]
                last -= 1
                while nums[last] == val and last > 0:
                    last -= 1
                    
            fast += 1
            
        return last + 1