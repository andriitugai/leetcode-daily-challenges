class Solution:
    def tupleSameProduct(self, nums: List[int]) -> int:
        counter = defaultdict(int)
        result = 0

        for i in range(len(nums)):
            for j in range(i):
                prod = nums[i] * nums[j]
                result += counter[prod] * 8
                counter[prod] += 1

        return result