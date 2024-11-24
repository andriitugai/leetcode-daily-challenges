class Solution:
    def canAliceWin(self, n: int) -> bool:
        toRemove = 10
        round = 0
        while toRemove <= n:
            n -= toRemove
            round += 1
            toRemove -= 1
        return round % 2 == 1