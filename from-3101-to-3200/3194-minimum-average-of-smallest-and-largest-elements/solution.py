class Solution:
    def minimumAverage(self, nums: List[int]) -> float:
        nums.sort()
        n = len(nums)
        l, r = 0, n - 1
        minAvg = float("inf")
        while l < r:
            avg = float(nums[l] + nums[r])/2.0
            if avg < minAvg:
                minAvg = avg
            l += 1
            r -= 1

        return minAvg