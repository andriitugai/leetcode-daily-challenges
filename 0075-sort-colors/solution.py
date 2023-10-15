class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if not nums:
            return nums
        
        c = Counter(nums)
        i = 0
        for k in (0, 1, 2):
            if k in c:
                for _ in range(c[k]):
                    nums[i] = k
                    i += 1
                    