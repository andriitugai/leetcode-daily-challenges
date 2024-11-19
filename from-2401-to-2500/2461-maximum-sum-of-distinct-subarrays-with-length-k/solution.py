class Solution:
    def maximumSubarraySum(self, nums: List[int], k: int) -> int:
        result, cur_sum = 0, 0
        count = defaultdict(int)
        l = 0

        for r in range(len(nums)):
            cur_sum += nums[r]
            count[nums[r]] += 1

            if r - l + 1 > k:
                cur_sum -= nums[l]
                count[nums[l]] -= 1
                if count[nums[l]] == 0:
                    count.pop(nums[l])
                l += 1

            if r - l + 1 == k and len(count) == k :
                result = max(result, cur_sum)

        return result