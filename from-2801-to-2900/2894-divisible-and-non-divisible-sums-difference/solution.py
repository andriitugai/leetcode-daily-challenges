class Solution:
    def differenceOfSums(self, n: int, m: int) -> int:
        result = 0
        for num in range(1, n + 1):
            if num % m:
                result += num
            else:
                result -= num

        return result