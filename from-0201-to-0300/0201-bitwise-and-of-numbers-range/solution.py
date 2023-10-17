class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        p2 = 0
        while left != right:
            left >>= 1
            right >>= 1
            p2 += 1

        return right << p2