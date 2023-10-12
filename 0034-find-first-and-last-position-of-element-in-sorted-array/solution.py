class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        left = self.binarySearch(nums, target, True)
        right = self.binarySearch(nums, target, False)
        return [left, right]
    
        
    def binarySearch(self, nums, target, is_leftmost: bool) -> int:
        l, r = 0, len(nums) - 1
        ans = -1
        
        while l <= r:
            m = (l + r) // 2
            if nums[m] > target:
                r = m - 1
            elif nums[m] < target:
                l = m + 1
            else:
                ans = m
                if is_leftmost:
                    r = m - 1
                else:
                    l = m + 1
                    
        return ans