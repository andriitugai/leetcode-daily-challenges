class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        pre = [1] * len(nums)
        suf = [1] * len(nums)

        cur = 1
        for i, num in enumerate(nums):
            pre[i] = cur
            cur *= num

        cur = 1
        for i in range(len(nums)-1, -1, -1):
            suf[i] = cur
            cur *= nums[i]

        return [p * s for p, s in zip(pre, suf)]