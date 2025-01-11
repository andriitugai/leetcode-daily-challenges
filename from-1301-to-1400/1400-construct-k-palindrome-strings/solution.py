class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        n = len(s)

        if n < k:
            return False
        elif n == k:
            return True

        cnt = Counter(s)
        num_odds = 0
        for _, v in cnt.items():
            if v % 2 == 1:
                num_odds += 1
        if num_odds > k:
            return False

        return True