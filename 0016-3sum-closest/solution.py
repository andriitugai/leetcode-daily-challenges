lass Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans = sum(nums[:3])
        curdelta = abs(target-ans)
        
        n = len(nums)
        for i in range(n - 2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
                
            left, right = i + 1, n - 1
            
            while left < right:
                cursum = nums[i] + nums[left]  + nums[right]
                if cursum == target:
                    return cursum
                if abs(target-cursum) < curdelta:
                    ans = cursum
                    curdelta = abs(target-cursum)
                    
                if cursum < target:
                    left += 1
                else:
                    right -= 1
                    
        return ans
