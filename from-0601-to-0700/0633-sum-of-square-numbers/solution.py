class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        left = 0
        right = int(c ** 0.5)
        while left <= right:
            res = left*left + right*right
            if res == c:
                return True
            if res < c:
                left += 1
            else:
                right -= 1

        return False