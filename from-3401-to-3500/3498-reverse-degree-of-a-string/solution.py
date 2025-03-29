class Solution:
    def reverseDegree(self, s: str) -> int:
        result = 0
        for i, c in enumerate(s, start=1):
            result += i * (26 - ord(c) + ord('a'))
        return result
        