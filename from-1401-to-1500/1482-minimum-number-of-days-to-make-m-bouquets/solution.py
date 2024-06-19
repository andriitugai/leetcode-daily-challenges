class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        if len(bloomDay) < m * k:
            return -1

        def is_possible(d: int) -> bool:
            nonlocal m
            adj = 0
            bouq = 0
            for b in bloomDay:
                if b > d:
                    adj = 0
                else:
                    adj += 1
                    if adj == k:
                        bouq += 1
                        adj = 0
            return bouq >= m

        left = min(bloomDay)
        right = max(bloomDay)

        while left < right:
            mid  = left + (right - left) // 2
            if not is_possible(mid):
                left = mid + 1
            else:
                right = mid

        return left