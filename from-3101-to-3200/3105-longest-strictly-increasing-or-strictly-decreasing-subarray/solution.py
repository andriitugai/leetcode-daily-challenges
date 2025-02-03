class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        incr, decr = 1, 1
        prev = nums[0]
        incr_max, decr_max = 0, 0
        for num in nums[1:]:
            if num > prev:
                decr_max = max(decr_max, decr)
                incr += 1
                decr = 1
            elif num < prev:
                incr_max = max(incr_max, incr)
                incr = 1
                decr += 1
            else:
                incr_max = max(incr_max, incr)
                decr_max = max(decr_max, decr)
                incr, decr = 1, 1

            prev = num
            
        return max(max(incr_max, incr), max(decr_max, decr))