class Solution:
    def maximumDifference(self, nums: List[int]) -> int:
        cur_min = float('inf')
        max_diff = -1

        for num in nums:
            if num < cur_min:
                cur_min = num
            elif num - cur_min > max_diff:
                max_diff = num - cur_min

        if max_diff == 0:
            max_diff = -1

        return max_diff