class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        sum1, sum2 = sum(nums1), sum(nums2)
        zer1, zer2 = sum(1 for num in nums1 if num == 0), sum(1 for num in nums2 if num == 0)

        if zer1 == 0 and zer2 == 0:
            return sum1 if sum1 == sum2 else -1
        if zer1 == 0:
            return sum1 if sum2 + zer2 <= sum1 else -1
        if zer2 == 0:
            return sum2 if sum1 + zer1 <= sum2 else -1

        return max(sum1 + zer1, sum2 + zer2)