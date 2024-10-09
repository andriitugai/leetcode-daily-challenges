class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        open = 0
        result = 0
        for c in s:
            if c == "(":
                open += 1
            elif open > 0:
                open -= 1
            elif open == 0:
                result += 1

        return result + open