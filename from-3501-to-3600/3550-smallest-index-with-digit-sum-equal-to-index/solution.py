class Solution:
    def smallestIndex(self, nums: List[int]) -> int:
        def sumOfDigits(n: int) -> int:
            s = 0
            while n > 0:
                s += n % 10
                n //= 10
            return s

        for i, num in enumerate(nums):
            if i == sumOfDigits(num):
                return i

        return -1