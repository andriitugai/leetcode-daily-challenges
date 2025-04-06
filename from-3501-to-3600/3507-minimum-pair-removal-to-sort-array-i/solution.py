class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        result = 0
        
        while not all(nums[i] >= nums[i - 1] for i in range(1, len(nums))):
            min_sum = 5000
            min_idx = -1
            for i in range(len(nums) - 1):
                pair_sum = nums[i] + nums[i + 1]
                if pair_sum < min_sum:
                    min_sum = pair_sum
                    min_idx = i

            nums[min_idx] = min_sum
            nums.pop(min_idx + 1)
            result += 1

        return result
