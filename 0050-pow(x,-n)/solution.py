class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 1:
            return x
        elif n == 0:
            return 1
        else:
            if n < 0:
                return 1 / self.myPow(x, -n)
            else:
                if n % 2:
                    return x * self.myPow(x, n - 1)
        return self.myPow(x * x, n // 2)