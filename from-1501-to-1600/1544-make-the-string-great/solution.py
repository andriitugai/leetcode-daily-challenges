class Solution:
    def makeGood(self, s: str) -> str:
        result = []
        for c in s:
            if result and c != result[-1] and c.lower() == result[-1].lower():
                result.pop()
            else:
                result.append(c)

        return ''.join(result)