class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        def bin_search(l, r, target) -> int:
            # Return largest i where nums[i] < target
            while l <= r:
                m = (l + r) // 2
                if nums[m] >= target:
                    r = m - 1
                else:
                    l = m + 1
            return r

        nums.sort()
        result = 0
        n = len(nums)
        for i in range(n):
            low = lower - nums[i]
            up = upper - nums[i]
            result += (
                bin_search(i + 1, n - 1, up + 1) - 
                bin_search(i + 1, n - 1, low)
            )
        return result
