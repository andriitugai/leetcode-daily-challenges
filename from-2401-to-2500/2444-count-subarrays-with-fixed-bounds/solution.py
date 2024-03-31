class Solution:
    def countSubarrays(self, nums: List[int], minK: int, maxK: int) -> int:
        bad = mi = ma = -1
        result = 0

        for i, n in enumerate(nums):
            if not minK <= n <= maxK:
                bad = i

            if minK == n:
                mi = i
            if maxK == n:
                ma = i

            start = min(mi, ma)
            if start > bad:
                result += (start - bad)

        return result
        