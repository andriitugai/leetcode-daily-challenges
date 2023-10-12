class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        def twoSum(nums, target):

            if len(nums) < 2:
                return None

            result = {}
            additions = {}

            for num in nums:
                if target - num in additions:
                    result[target - num] = additions[target - num]
                else:
                    additions[num] = target - num

            return result

        nums = sorted(nums)

        result = []
        prev = object()

        for idx, num in enumerate(nums):
            if num != prev:
                pairs = twoSum(nums[idx + 1:], 0 - num)
                if pairs:
                    result.extend([[num, a, b] for a, b in pairs.items()])

            prev = num

        return result

