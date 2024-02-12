class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        for key, val in cnt.items():
            if val > len(nums) // 2:
                return key

        return -1