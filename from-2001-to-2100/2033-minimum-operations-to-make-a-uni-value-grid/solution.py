class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        nums = [ num for row in grid for num in row ]
        for num in nums:
            if (num - nums[0]) % x != 0:
                return -1

        nums.sort()
        med = nums[len(nums) // 2]

        return sum(abs(num - med) // x for num in nums)