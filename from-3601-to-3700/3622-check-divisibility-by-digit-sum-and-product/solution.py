class Solution:
    def checkDivisibility(self, n: int) -> bool:
        num = n
        dSumm, dProd = 0, 1
        while n > 0:
            dig = n % 10
            dSumm += dig
            dProd *= dig
            n //= 10

        return num % (dSumm + dProd) == 0
        