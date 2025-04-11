class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        result = 0
        for num in range(11, 100, 11):
            if low <= num <= high:
                result += 1

        for num in range(1000, high + 1):
            if num >= low:
                if (num // 1000) + ((num // 100) % 10) == num % 10 + (num % 100) // 10:
                    result += 1

        return result