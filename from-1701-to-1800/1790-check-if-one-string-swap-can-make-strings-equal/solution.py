class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        if Counter(s1) != Counter(s2):
            return False

        half = 0
        for c1, c2 in zip(s1, s2):
            if c1 != c2:
                half += 1

        return half == 0 or half == 2