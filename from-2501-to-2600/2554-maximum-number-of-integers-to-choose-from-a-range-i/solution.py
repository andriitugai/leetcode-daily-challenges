class Solution:
    def maxCount(self, banned: List[int], n: int, maxSum: int) -> int:
        isBanned = defaultdict(bool)
        for b in banned:
            if b <= n:
                isBanned[b] = True

        sum, cnt = 0, 0
        for i in range(1, n + 1):
            if not isBanned[i] and i + sum <= maxSum:
                cnt += 1
                sum += i

        return cnt