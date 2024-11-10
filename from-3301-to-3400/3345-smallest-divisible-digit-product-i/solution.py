class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        def getProduct(num: int) -> int:
            product = 1
            while num > 1:
                product *= (num % 10)
                num //= 10
            return product

        for c in range(t):
            if getProduct(n + c) % t == 0:
                return n + c

        return -1