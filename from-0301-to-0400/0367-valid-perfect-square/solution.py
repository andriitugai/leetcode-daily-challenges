class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True

        left, right = 1, num // 2
        while left <= right:
            mid = (left + right) // 2
            m2 = mid * mid
            if m2 == num:
                return True
            elif m2 > num:
                right = mid - 1
            else:
                left = mid + 1

        return False 