class Solution:
    def concatHex36(self, n: int) -> str:
        def toBase(x: int, b: int) -> str:
            symb = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'
            result = ''
            while x > 0:
                result = symb[x % b] + result
                x //= b
            return result

        return toBase(n * n, 16) + toBase(n * n * n, 36)