class Solution:
    def isArraySpecial(self, nums: List[int], queries: List[List[int]]) -> List[bool]:
        result = [False] * len(queries)
        preSum = [0] * len(nums)
        preSum[0] = 1
        for i in range(1, len(nums)):
            preSum[i] = preSum[i - 1]
            if nums[i] % 2 != nums[i - 1] % 2:
                preSum[i] += 1

        for i, q in enumerate(queries):
            s, e = q
            if preSum[e] - preSum[s] == e - s:
                result[i] = True
        return result
        