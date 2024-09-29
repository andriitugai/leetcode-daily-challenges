class Solution:
    def minElement(self, nums: List[int]) -> int:
        minSum = 1000
        for num in nums:
            sum = 0
            while num > 0:
                sum += num % 10
                num //= 10
            if sum < minSum:
                minSum = sum

        return minSum
        