class Solution:
    def check(self, nums: List[int]) -> bool:
        n = len(nums)
        rotation_points = 0
        for i in range(1, n):
            if nums[i - 1] > nums[i]:
                rotation_points += 1
                
        if nums[-1] > nums[0]:
            rotation_points += 1
        
        return rotation_points <= 1